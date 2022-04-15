package main

import (
	"augments/models"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"runtime/debug"
	"strconv"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s/n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter) {
	trace := fmt.Sprintf("%s/n%s", http.StatusText(http.StatusBadRequest), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (app *application) writeStruct(w http.ResponseWriter, v interface{}) {
	resp, err := json.Marshal(v)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func getCredentials(r *http.Request) (string, uint64) {
	// Format
	// X-Aug-ID: xxx
	// X-Aug-Token: xxx
	userID, err := strconv.ParseUint(r.Header.Get("X-Aug-ID"), 10, 32)
	if err != nil || userID == 0 {
		return "", 0
	}
	token := r.Header.Get("X-Aug-Token")
	return token, userID
}

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
