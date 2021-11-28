package main

import (
	"augments/models"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s/n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) returnStruct(w http.ResponseWriter, v interface{}) {
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

func generateAccessToken() (string, error) {
	// Make a 256-bit string
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
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

func (app *application) createSteamUser(steamID uint64) error {
	// Start a transaction
	tx, err := app.db.Beginx()
	if err != nil {
		return err
	}

	// Create user
	_, err = tx.Exec("INSERT INTO user (steam_id, egg_cap, storage_pages) VALUES (?, ?, ?)", steamID, 5, 1)
	if err != nil {
		tx.Rollback()
		return err
	}

	user := &models.User{}
	err = tx.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert initial staff
	_, err = tx.Exec("INSERT INTO staff (user_id, slot, name) VALUES (?, ?, ?)", user.ID, 0, "STAFF 1")
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert initial creatures
	species := &[]models.Species{}
	err = tx.Select(species, "SELECT * FROM species WHERE id != 0 ORDER BY RAND() LIMIT 5")
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, s := range *species {
		// Initial eggs are ready to hatch
		neededWins := s.GetRequiredEggWins()

		// Insert creatures
		_, err := tx.Exec(
			"INSERT INTO creature (user_id, species_id, name, is_egg, staff_slot, wins) VALUES (?, ?, ?, ?, ?, ?)",
			user.ID, s.ID, s.Name, true, -1, neededWins)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Insert initial 28 actions
	for i := 15; i <= 42; i++ {
		_, err = tx.Exec("INSERT INTO user_action (user_id, action_id, qty) VALUES (?, ?, ?)", user.ID, i, 1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
