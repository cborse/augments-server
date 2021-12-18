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
		app.clientError(w, http.StatusUnauthorized)
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

	// Select all skills
	skills := &[]models.Skill{}
	err = app.db.Select(skills, "SELECT * FROM skill ORDER BY id")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Select all skillsets
	skillsets := &[]models.Skillset{}
	err = app.db.Select(skillsets, "SELECT * FROM skillset")
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
	err = app.db.Select(staffs, "SELECT * FROM staff WHERE user_id = ? ORDER BY id", user.ID)
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

	// Get all user actions
	userActions := &[]models.UserAction{}
	err = app.db.Select(userActions, "SELECT * FROM user_action WHERE user_id = ? AND qty > 0", user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Get all user skills
	userSkills := &[]models.UserSkill{}
	err = app.db.Select(userSkills, "SELECT * FROM user_skill WHERE user_id = ? AND qty > 0", user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Return the struct
	data := struct {
		Actions     *[]models.Action     `json:"actions"`
		Actionsets  *[]models.Actionset  `json:"actionsets"`
		Creatures   *[]models.Creature   `json:"creatures"`
		Series      *[]models.Series     `json:"series"`
		Skills      *[]models.Skill      `json:"skills"`
		Skillsets   *[]models.Skillset   `json:"skillsets"`
		Species     *[]models.Species    `json:"species"`
		Staffs      *[]models.Staff      `json:"staffs"`
		User        *models.User         `json:"user"`
		UserActions *[]models.UserAction `json:"user_actions"`
		UserSkills  *[]models.UserSkill  `json:"user_skills"`
	}{
		Actions:     actions,
		Actionsets:  actionsets,
		Creatures:   creatures,
		Series:      series,
		Skills:      skills,
		Skillsets:   skillsets,
		Species:     species,
		Staffs:      staffs,
		User:        user,
		UserActions: userActions,
		UserSkills:  userSkills,
	}

	app.returnStruct(w, data)
}

func (app *application) assign(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		StaffID    uint64 `json:"staff_id"`
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
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	// Make sure there's enough room
	count := 0
	err = app.db.Get(&count, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND staff_id = ?", userID, body.StaffID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if count >= 8 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Assign the creature to the staff
	_, err = app.db.Exec("UPDATE creature SET staff_id = ? WHERE id = ?", body.StaffID, body.CreatureID)
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
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	// Unassign the creature from the staff
	_, err = app.db.Exec("UPDATE creature SET staff_id = 0 WHERE id = ?", body.CreatureID)
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
		app.clientError(w, http.StatusUnauthorized)
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
		app.clientError(w, http.StatusForbidden)
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
	err = app.db.Get(&storageCount, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND egg = false", userID)
	if err != nil {
		app.serverError(w, err)
		return
	} else if storageCount >= uint(user.StoragePages*20) {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Hatch it
	_, err = app.db.Exec("UPDATE creature SET egg = false, wins = 0 WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// OK!
}

func (app *application) replaceAction(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		ActionID   uint32 `json:"action_id"`
		Slot       uint8  `json:"slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the requester's user ID
	_, userID := getCredentials(r)

	// Valid slot
	if body.Slot > 2 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// User ID
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if creature.UserID != userID {
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	// In inventory and qty > 0
	userAction := &models.UserAction{}
	err = app.db.Get(userAction, "SELECT * FROM user_action WHERE user_id = ? AND action_id = ?", userID, body.ActionID)
	if err != nil {
		app.serverError(w, err)
		return
	} else if userAction.Qty == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Not egg
	if creature.Egg {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Creature doesn't already know this action
	if creature.Action1 == body.ActionID || creature.Action2 == body.ActionID || creature.Action3 == body.ActionID {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Species can learn this action
	action := &models.Action{}
	err = app.db.Get(action, "SELECT * FROM action WHERE id = ?", body.ActionID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	species := &models.Species{}
	err = app.db.Get(species, "SELECT * FROM species WHERE id = ?", creature.SpeciesID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Core
	if action.Core && species.Type1 != action.Type && species.Type2 != action.Type && species.Type3 != action.Type {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Actionset
	actionset := &models.Actionset{}
	err = app.db.Get(actionset, "SELECT * FROM actionset WHERE species_id = ? AND action_id = ? AND series_id = ?", creature.SpeciesID, body.ActionID, creature.SeriesID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Start a transaction
	tx, err := app.db.Beginx()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove from inventory
	_, err = tx.Exec("UPDATE user_action SET qty = qty - 1 WHERE user_id = ? AND action_id = ?", userID, body.ActionID)
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Set on creature
	if body.Slot == 0 {
		_, err = tx.Exec("UPDATE creature SET action1 = ? WHERE id = ?", body.ActionID, body.CreatureID)
	} else if body.Slot == 1 {
		_, err = tx.Exec("UPDATE creature SET action2 = ? WHERE id = ?", body.ActionID, body.CreatureID)
	} else if body.Slot == 2 {
		_, err = tx.Exec("UPDATE creature SET action3 = ? WHERE id = ?", body.ActionID, body.CreatureID)
	} else {
		tx.Rollback()
		app.clientError(w, http.StatusBadRequest)
		return
	}
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Should be OK
	tx.Commit()
}
