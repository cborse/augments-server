package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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

func (app *application) createSteamUser(steamID uint64) (*User, error) {
	// Start a transaction
	tx, err := app.db.Beginx()
	if err != nil {
		return nil, err
	}

	// Create user and select him
	err = User_create(tx, steamID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user, err := User_getBySteamID(tx, steamID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create initial staff
	err = Staff_create(tx, user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create initial creatures
	// - first, select 5 random species
	// - second, create the creatures
	species, err := Species_getRandom(tx, 5)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, species := range *species {
		err = Creature_create(tx, user.ID, &species, 1, true)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Create initial 28 user actions
	for i := uint32(15); i <= 42; i++ {
		err = UserAction_add(tx, user.ID, i)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Create initial 33 user skills
	for i := uint32(1); i <= 33; i++ {
		err = UserSkill_add(tx, user.ID, i)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

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
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	// Find the user
	user, err := User_getBySteamID(app.db, steamID)
	if err == sql.ErrNoRows {
		// User doesn't exist; create one
		user, err = app.createSteamUser(steamID)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	// Assign a new token
	if err = User_assignNewToken(app.db, user); err != nil {
		app.serverError(w, err)
		return
	}

	// Return the token
	response := struct {
		ID    uint64 `json:"id"`
		Token string `json:"token"`
	}{
		ID:    user.ID,
		Token: user.Token,
	}
	app.returnStruct(w, response)
}

func (app *application) getData(w http.ResponseWriter, r *http.Request) {
	_, userID := getCredentials(r)
	user, err := User_get(app.db, userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	creatures, err := Creature_select(app.db, userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	staffs, err := Staff_select(app.db, userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	userActions, err := UserAction_select(app.db, userID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	}

	userSkills, err := UserSkill_select(app.db, userID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	}

	actions, err := Action_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	actionsets, err := Actionset_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	series, err := Series_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	skills, err := Skill_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	skillsets, err := Skillset_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	species, err := Species_select(app.db)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Return the struct
	data := struct {
		Actions     *[]Action     `json:"actions"`
		Actionsets  *[]Actionset  `json:"actionsets"`
		Creatures   *[]Creature   `json:"creatures"`
		Series      *[]Series     `json:"series"`
		Skills      *[]Skill      `json:"skills"`
		Skillsets   *[]Skillset   `json:"skillsets"`
		Species     *[]Species    `json:"species"`
		Staffs      *[]Staff      `json:"staffs"`
		User        *User         `json:"user"`
		UserActions *[]UserAction `json:"user_actions"`
		UserSkills  *[]UserSkill  `json:"user_skills"`
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
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		StaffSlot  uint8  `json:"staff_slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	creature, err := Creature_get(app.db, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure there's room on the new staff
	count, err := Creature_getStaffCount(app.db, userID, body.StaffSlot)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if count >= 8 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Good to go
	if err = Creature_updateStaffSlot(app.db, creature.ID, int8(body.StaffSlot)); err != nil {
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

	creature, err := Creature_get(app.db, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	// Good to go
	if err = Creature_updateStaffSlot(app.db, creature.ID, -1); err != nil {
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

	creature, err := Creature_get(app.db, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure it has the required wins
	species, err := Species_get(app.db, creature.SpeciesID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	reqWins := Species_getRequiredEggWins(species)
	if creature.Wins < uint32(reqWins) {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure there's enough room in storage
	user, err := User_get(app.db, userID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	count, err := Creature_getStorageCount(app.db, userID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if count >= int(user.StoragePages*20) {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Good to go
	if err = Creature_hatch(app.db, creature.ID); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) learnAction(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		ActionID   uint32 `json:"action_id"`
		Slot       uint8  `json:"slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	creature, err := Creature_get(app.db, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature is not an egg
	if creature.Egg {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature doesn't already know this action
	if creature.Action1 == body.ActionID || creature.Action2 == body.ActionID || creature.Action3 == body.ActionID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature can learn this action
	species, err := Species_get(app.db, creature.SpeciesID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	action, err := Action_get(app.db, body.ActionID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if action.Core && species.Type1 != action.Type && species.Type2 != action.Type && species.Type3 != action.Type {
		canLearn, err := Actionset_exists(app.db, creature.SpeciesID, body.ActionID, creature.SeriesID)
		if err != nil {
			app.serverError(w, err)
			return
		}
		if !canLearn {
			app.clientError(w, http.StatusForbidden)
			return
		}
	}

	// Make sure the user owns this action
	userAction, err := userAction_get(app.db, userID, body.ActionID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if err == sql.ErrNoRows || userAction.Qty == 0 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Good to go
	tx, err := app.db.Beginx()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove from inventory
	if err = UserAction_remove(tx, userAction); err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Set on creature
	if err = Creature_learnAction(tx, creature.ID, body.ActionID, body.Slot); err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Should be OK
	tx.Commit()
}

func (app *application) learnSkill(w http.ResponseWriter, r *http.Request) {
	body := struct {
		CreatureID uint64 `json:"creature_id"`
		SkillID    uint32 `json:"skill_id"`
		Slot       uint8  `json:"slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure it's a valid slot
	if body.Slot > 2 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	creature, err := Creature_get(app.db, body.CreatureID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Make sure the user IDs match
	_, userID := getCredentials(r)
	if creature.UserID != userID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature is not an egg
	if creature.Egg {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature doesn't already know this skill
	if creature.Skill1 == body.SkillID || creature.Skill2 == body.SkillID || creature.Skill3 == body.SkillID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Make sure creature can learn this skill
	skill, err := Skill_get(app.db, body.SkillID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if !skill.Core {
		canLearn, err := Skillset_exists(app.db, creature.SpeciesID, body.SkillID)
		if err != nil {
			app.serverError(w, err)
			return
		}
		if !canLearn {
			app.clientError(w, http.StatusForbidden)
			return
		}
	}

	// Make sure the user owns this skill
	userSkill, err := UserSkill_get(app.db, userID, body.SkillID)
	if err != nil && err != sql.ErrNoRows {
		app.serverError(w, err)
		return
	} else if err == sql.ErrNoRows || userSkill.Qty == 0 {
		app.clientError(w, http.StatusForbidden)
		return
	}

	// Good to go
	tx, err := app.db.Beginx()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove from inventory
	if err = UserSkill_remove(tx, userSkill); err != nil {
		tx.Rollback()
		app.serverError(w, err)
		return
	}

	// Set on creature
	if err = Creature_learnSkill(tx, creature.ID, body.SkillID, body.Slot); err != nil {
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
