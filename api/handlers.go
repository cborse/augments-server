package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"augments/models"
)

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		Ticket  string `json:"ticket"`
		SteamID uint64 `json:"steam_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Verify with Steam servers
	steamID, err := authenticateSteamTicket(body.Ticket)
	if err != nil {
		app.serverError(w, err)
		return
	} else if steamID != body.SteamID {
		clientError(w, http.StatusUnauthorized)
		return
	}

	// Find the user
	user := &models.User{}
	err = app.db.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
	if err == sql.ErrNoRows {
		// User doesn't exist; create one
		if err = app.createSteamUser(steamID); err != nil {
			app.serverError(w, err)
			return
		}

		// Find user now
		err = app.db.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	// Generate a new token
	user.Token, err = generateAccessToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Update user
	_, err = app.db.Exec("UPDATE user SET token = ? WHERE id = ?", user.Token, user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Select all actions
	actions := &[]models.Action{}
	err = app.db.Select(actions, "SELECT * FROM action ORDER BY id")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Select all actionsets
	actionsets := &[]models.Actionset{}
	err = app.db.Select(actionsets, "SELECT * FROM actionset")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Select all series
	series := &[]models.Series{}
	err = app.db.Select(series, "SELECT * FROM series ORDER BY id")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Select all species
	species := &[]models.Species{}
	err = app.db.Select(species, "SELECT * FROM species ORDER BY id")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Get the staffs
	staffs := &[]models.Staff{}
	err = app.db.Select(staffs, "SELECT * FROM staff WHERE user_id = ? ORDER BY slot", user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Get the creatures
	creatures := &[]models.Creature{}
	err = app.db.Select(creatures, "SELECT * FROM creature WHERE user_id = ? ORDER BY id", user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Return the struct
	data := struct {
		Actions    *[]models.Action    `json:"actions"`
		Actionsets *[]models.Actionset `json:"actionsets"`
		Creatures  *[]models.Creature  `json:"creatures"`
		Series     *[]models.Series    `json:"series"`
		Species    *[]models.Species   `json:"species"`
		Staffs     *[]models.Staff     `json:"staffs"`
		User       *models.User        `json:"user"`
	}{
		Actions:    actions,
		Actionsets: actionsets,
		Creatures:  creatures,
		Series:     series,
		Species:    species,
		Staffs:     staffs,
		User:       user,
	}

	app.returnStruct(w, data)
}

func (app *application) assign(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		StaffSlot  int8   `json:"staff_slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the requester's user ID
	_, userID := getCredentials(r)

	// Find the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	if creature.UserID != userID {
		clientError(w, http.StatusUnauthorized)
		return
	}

	// Make sure there's enough room
	count := 0
	err = app.db.Get(&count, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND staff_slot = ?", userID, body.StaffSlot)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if count >= 8 {
		clientError(w, http.StatusForbidden)
		return
	}

	// Assign the creature to the staff
	_, err = app.db.Exec("UPDATE creature SET staff_slot = ? WHERE id = ?", body.StaffSlot, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// OK!
}

func (app *application) unassign(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		CreatureID uint64 `json:"creature_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the requester's user ID
	_, userID := getCredentials(r)

	// Find the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	if creature.UserID != userID {
		clientError(w, http.StatusUnauthorized)
		return
	}

	// Unassign the creature from the staff
	_, err = app.db.Exec("UPDATE creature SET staff_slot = -1 WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// OK!
}

func (app *application) hatchEgg(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		CreatureID uint64 `json:"creature_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the requester's user ID
	_, userID := getCredentials(r)

	// Find the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user ID's match
	if creature.UserID != userID {
		clientError(w, http.StatusUnauthorized)
		return
	}

	// Make sure it has the required wins
	species := &models.Species{}
	err = app.db.Get(species, "SELECT * FROM species WHERE id = ?", creature.SpeciesID)
	if err != nil {
		app.serverError(w, err)
		return
	} else if species.Rarity == 0 && creature.Wins < 2 ||
		species.Rarity == 1 && creature.Wins < 4 ||
		species.Rarity == 2 && creature.Wins < 8 {
		clientError(w, http.StatusForbidden)
		return
	}

	// Make sure there's enough room in storage
	user := &models.User{}
	err = app.db.Get(user, "SELECT * FROM user WHERE id = ?", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	var storageCount uint
	err = app.db.Get(&storageCount, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND is_egg = false", userID)
	if err != nil {
		app.serverError(w, err)
		return
	} else if storageCount >= uint(user.StoragePages*20) {
		clientError(w, http.StatusForbidden)
		return
	}

	// Hatch it
	_, err = app.db.Exec("UPDATE creature SET is_egg = false, wins = 0 WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// OK!
}
