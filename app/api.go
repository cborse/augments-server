package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"augments/models"
)

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	body := LoginRequest{}
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
	} else if err != nil {
		app.serverError(w, err)
		return
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
	response := LoginResponse{ID: user.ID, Token: token}
	app.writeStruct(w, response)
}

func (app *application) getData(w http.ResponseWriter, r *http.Request) {
	_, userID := getCredentials(r)

	user := &models.User{}
	err := app.db.Get(user, "SELECT * FROM user WHERE id = ?", userID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	staffs := &[]models.Staff{}
	err = app.db.Select(staffs, "SELECT * FROM staff WHERE user_id = ? ORDER BY slot", userID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	creatures := &[]models.Creature{}
	err = app.db.Select(creatures, "SELECT * FROM creature WHERE user_id = ? ORDER BY id", userID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
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
	data := GetDataResponse{
		User:        user,
		Staffs:      staffs,
		Creatures:   creatures,
		UserActions: userActions,
		UserSkills:  userSkills,
	}
	app.writeStruct(w, data)
}

func (app *application) assign(w http.ResponseWriter, r *http.Request) {
	body := AssignRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w)
		return
	}

	// Make sure staff exists
	tmp := 0
	err = app.db.Get(&tmp, "SELECT COUNT(*) FROM staff WHERE user_id = ? AND slot = ?", userID, body.StaffSlot)
	if err != nil {
		app.serverError(w, err)
		return
	} else if tmp != 1 {
		app.clientError(w)
		return
	}

	// Make sure there's room on the new staff
	err = app.db.Get(&tmp, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND staff_slot = ?", userID, body.StaffSlot)
	if err != nil {
		app.serverError(w, err)
		return
	} else if tmp >= 8 {
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
	body := UnassignRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
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
	body := HatchEggRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Get the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's an egg
	if !creature.Egg {
		app.clientError(w)
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
	body := LearnActionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w)
		return
	}

	// Make sure it's a valid action
	if body.ActionID <= models.ACTION_NONE || body.ActionID >= models.ACTION_COUNT {
		app.clientError(w)
		return
	}

	// Get the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's not an egg
	if creature.Egg {
		app.clientError(w)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
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
	body := LearnSkillRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w)
		return
	}

	// Make sure it's a valid skill
	if body.SkillID <= models.SKILL_NONE || body.SkillID >= models.SKILL_COUNT {
		app.clientError(w)
		return
	}

	// Get the creature
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", body.CreatureID)
	if err == sql.ErrNoRows {
		app.clientError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's not an egg
	if creature.Egg {
		app.clientError(w)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
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

func (app *application) matchmake(w http.ResponseWriter, r *http.Request) {
	body := MatchmakeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	_, userID := getCredentials(r)

	// Find the highest level creature on this staff
	creatures := []models.Creature{}
	err := app.db.Select(&creatures, "SELECT * FROM creature WHERE user_id = ? AND staff_slot = ? ORDER BY wins DESC", userID, body.StaffSlot)
	if err != nil || len(creatures) < 5 {
		app.serverError(w, err)
		return
	}
	highestLevel := creatures[0].GetLevel()

	// Find a lobby
	lobby := app.matchMaker.findOrCreateLobby(userID, highestLevel, body.StaffSlot)

	// If the lobby was created (this user is host), wait for a match
	if !lobby.ready {
		app.matchMaker.waitForMatch(lobby)
	}
}
