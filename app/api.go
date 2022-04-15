package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"augments/models"
)

func authenticateSteamTicket(ticket string) (uint64, error) {
	// Verify with Steam servers
	url := "https://partner.steam-api.com/ISteamUserAuth/AuthenticateUserTicket/v1/"
	url += "?key=DD6091EF398948CC15AEE5F381D08D6B"
	url += "&appid=1390280"
	url += "&ticket=" + ticket

	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return 0, err
	}

	steamBody := struct {
		Response struct {
			Params struct {
				Result          string `json:"result"`
				SteamID         string `json:"steamid"`
				OwnerSteamID    string `json:"ownersteamid"`
				VACBanned       bool   `json:"vacbanned"`
				PublisherBanned bool   `json:"publisherbanned"`
			} `json:"params"`
		} `json:"response"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&steamBody)
	if err != nil {
		return 0, err
	}

	if steamBody.Response.Params.Result != "OK" || steamBody.Response.Params.VACBanned || steamBody.Response.Params.PublisherBanned {
		return 0, nil
	}

	return strconv.ParseUint(steamBody.Response.Params.SteamID, 10, 64)
}

func (app *application) createSteamUser(steamID uint64) (*models.User, error) {
	// Start a transaction
	tx, err := app.db.Beginx()
	if err != nil {
		return nil, err
	}

	// Create user and select him
	_, err = tx.Exec("INSERT INTO user (steam_id, egg_cap, storage_pages, credits) VALUES (?, 5, 1, 100)", steamID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user := &models.User{}
	err = tx.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create initial staff
	_, err = tx.Exec("INSERT INTO staff (user_id, slot, name) VALUES (?, 0, 'STAFF 1')", user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create initial creatures
	// - first, select 5 random species IDs
	// - second, insert the creatures in the database
	speciesIDs := []models.SpeciesID{}
	for len(speciesIDs) < 5 {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(models.SPECIES_COUNT-2)))
		if err != nil {
			return nil, err
		}
		speciesID := models.SpeciesID(n.Uint64()) + 1
		exists := false
		for _, id := range speciesIDs {
			if id == speciesID {
				exists = true
				break
			}
		}
		if !exists {
			speciesIDs = append(speciesIDs, speciesID)
		}
	}
	for _, speciesID := range speciesIDs {
		species := models.GetSpecies(speciesID)
		_, err := tx.Exec(
			"INSERT INTO creature (user_id, species_id, name, egg, staff_slot, wins, action1, action2) VALUES (?, ?, ?, true, -1, 8, ?, ?)",
			user.ID, speciesID, species.Name, species.Type1, species.Type2)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Insert initial 28 actions
	for i := 15; i <= 42; i++ {
		_, err = tx.Exec("INSERT INTO user_action (user_id, action_id, qty) VALUES (?, ?, 1)", user.ID, i)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Insert initial 32 skills
	for i := 1; i <= 33; i++ {
		_, err = tx.Exec("INSERT INTO user_skill (user_id, skill_id, qty) VALUES (?, ?, 1)", user.ID, i)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// OK
	return user, tx.Commit()
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	body := struct {
		SteamID     uint64 `json:"steam_id"`
		SteamTicket string `json:"steam_ticket"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Verify with Steam servers
	steamID, err := authenticateSteamTicket(body.SteamTicket)
	if err != nil {
		app.serverError(w, err)
		return
	} else if steamID != body.SteamID {
		app.clientError(w)
		return
	}

	// Find the user
	user := &models.User{}
	err = app.db.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
	if err == sql.ErrNoRows {
		// User doesn't exist; create one
		user, err = app.createSteamUser(steamID)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	// Create a new access token
	bytes := make([]byte, 32)
	_, err = rand.Read(bytes)
	if err != nil {
		app.serverError(w, err)
		return
	}
	token := base64.URLEncoding.EncodeToString(bytes)

	// Update the user
	_, err = app.db.Exec("UPDATE user SET token = ? WHERE id = ?", token, user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Return the user ID and token
	response := struct {
		ID    uint64 `json:"id"`
		Token string `json:"token"`
	}{
		ID:    user.ID,
		Token: token,
	}
	app.writeStruct(w, response)
}

func (app *application) getData(w http.ResponseWriter, r *http.Request) {
	_, userID := getCredentials(r)

	user := &models.User{}
	err := app.db.Get(user, "SELECT * FROM user WHERE id = ?", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	staffs := &[]models.Staff{}
	err = app.db.Select(staffs, "SELECT * FROM staff WHERE user_id = ? ORDER BY slot", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	creatures := &[]models.Creature{}
	err = app.db.Select(creatures, "SELECT * FROM creature WHERE user_id = ? ORDER BY id", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	userActions := &[]models.UserAction{}
	err = app.db.Select(userActions, "SELECT * FROM user_action WHERE user_id = ?", userID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	}

	userSkills := &[]models.UserSkill{}
	err = app.db.Select(userSkills, "SELECT * FROM user_skill WHERE user_id = ?", userID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	}

	// Return the struct
	data := struct {
		User        *models.User         `json:"user"`
		Staffs      *[]models.Staff      `json:"staffs"`
		Creatures   *[]models.Creature   `json:"creatures"`
		UserActions *[]models.UserAction `json:"user_actions"`
		UserSkills  *[]models.UserSkill  `json:"user_skills"`
	}{
		User:        user,
		Staffs:      staffs,
		Creatures:   creatures,
		UserActions: userActions,
		UserSkills:  userSkills,
	}

	app.writeStruct(w, data)
}

func (app *application) assign(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		StaffSlot  uint8  `json:"staff_slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Make sure there's room on the new staff
	staffCount := 0
	err = app.db.Get(&staffCount, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND staff_slot = ?", userID, body.StaffSlot)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if staffCount >= 8 {
		app.clientError(w)
		return
	}

	// Good to go
	_, err = app.db.Exec("UPDATE creature SET staff_slot = ? WHERE id = ?", body.StaffSlot, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) unassign(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64 `json:"creature_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Good to go
	_, err = app.db.Exec("UPDATE creature SET staff_slot = ? WHERE id = ?", -1, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) hatchEgg(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64 `json:"creature_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Make sure it has the required wins
	species := models.GetSpecies(creature.SpeciesID)
	reqWins := uint32(math.Pow(2, float64(species.Rarity)+1))
	if creature.Wins < reqWins {
		app.clientError(w)
		return
	}

	// Make sure there's enough room in storage
	user := &models.User{}
	err = app.db.Get(user, "SELECT * FROM user WHERE id = ?", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	storageCount := 0
	err = app.db.Get(&storageCount, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND egg = false", userID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if storageCount >= int(user.StoragePages*20) {
		app.clientError(w)
		return
	}

	// Good to go
	_, err = app.db.Exec("UPDATE creature SET egg = false, wins = 0 WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) learnAction(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64          `json:"creature_id"`
		ActionID   models.ActionID `json:"action_id"`
		Slot       uint8           `json:"slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w)
		return
	}

	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Make sure creature is not an egg
	if creature.Egg {
		app.clientError(w)
		return
	}

	// Make sure creature doesn't already know this action
	if creature.Action1 == body.ActionID || creature.Action2 == body.ActionID || creature.Action3 == body.ActionID {
		app.clientError(w)
		return
	}

	// Make sure creature can learn this action
	if !creature.CanLearnAction(body.ActionID) {
		app.clientError(w)
		return
	}

	// Make sure the user owns this action
	userAction := &models.UserAction{}
	err = app.db.Get(userAction, "SELECT * FROM user_action WHERE user_id = ? AND action_id = ?", userID, body.ActionID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if err == sql.ErrNoRows || userAction.Qty == 0 {
		app.clientError(w)
		return
	}

	// Good to go
	tx, err := app.db.Beginx()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove from inventory
	if userAction.Qty == 1 {
		_, err = tx.Exec("DELETE FROM user_action WHERE user_id = ? AND action_id = ?", userID, body.ActionID)
	} else {
		_, err = tx.Exec("UPDATE user_action SET qty = qty - 1 WHERE user_id = ? AND action_id = ?", userID, body.ActionID)
	}
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Set on creature
	_, err = tx.Exec("UPDATE creature SET action"+strconv.FormatUint(uint64(body.Slot+1), 10)+" = ? WHERE id = ?", body.ActionID, body.CreatureID)
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Should be OK
	if err = tx.Commit(); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) learnSkill(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64         `json:"creature_id"`
		SkillID    models.SkillID `json:"skill_id"`
		Slot       uint8          `json:"slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w)
		return
	}

	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Make sure creature is not an egg
	if creature.Egg {
		app.clientError(w)
		return
	}

	// Make sure creature doesn't already know this skill
	if creature.Skill1 == body.SkillID || creature.Skill2 == body.SkillID || creature.Skill3 == body.SkillID {
		app.clientError(w)
		return
	}

	// Make sure creature can learn this skill
	if !creature.CanLearnSkill(body.SkillID) {
		app.clientError(w)
		return
	}

	// Make sure the user owns this skill
	userSkill := &models.UserSkill{}
	err = app.db.Get(userSkill, "SELECT * FROM user_skill WHERE user_id = ? AND skill_id = ?", userID, body.SkillID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if err == sql.ErrNoRows || userSkill.Qty == 0 {
		app.clientError(w)
		return
	}

	// Good to go
	tx, err := app.db.Beginx()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove from inventory
	if userSkill.Qty == 1 {
		_, err = tx.Exec("DELETE FROM user_skill WHERE user_id = ? AND skill_id = ?", userID, body.SkillID)
	} else {
		_, err = tx.Exec("UPDATE user_skill SET qty = qty - 1 WHERE user_id = ? AND skill_id = ?", userID, body.SkillID)
	}
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Set on creature
	_, err = tx.Exec("UPDATE creature SET skill"+strconv.FormatUint(uint64(body.Slot+1), 10)+" = ? WHERE id = ?", body.SkillID, body.CreatureID)
	if err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Should be OK
	if err = tx.Commit(); err != nil {
		app.serverError(w, err)
		return
	}
}
