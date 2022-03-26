package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"augments/models"
)

type lobbyUser struct {
	id        uint64
	staffSlot uint8
	canceled  bool
}

type lobby struct {
	level int
	host  lobbyUser
	guest lobbyUser
}

type matchMaker struct {
	mutex   sync.Mutex
	lobbies []lobby
}

func (m *matchMaker) cleanLobbies() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i, lob := range m.lobbies {
		if lob.host.canceled && lob.guest.canceled {
			m.lobbies[i] = m.lobbies[len(m.lobbies)-1]
			m.lobbies = m.lobbies[:len(m.lobbies)-1]
		}
	}
}

func (m *matchMaker) findOrCreateLobby(id uint64, level int, staffSlot uint8) *lobby {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Check if there's a match
	for i, lob := range m.lobbies {
		if lob.guest.id == 0 && lob.level >= level-5 && lob.level <= level+5 {
			// Match found! Add the guest data
			m.lobbies[i].guest.id = id
			m.lobbies[i].guest.staffSlot = staffSlot
			return &m.lobbies[i]
		}
	}

	// No match. Create a lobby
	lobbyUser := lobbyUser{id: id, staffSlot: staffSlot}
	m.lobbies = append(m.lobbies, lobby{level: level, host: lobbyUser})

	return &m.lobbies[len(m.lobbies)-1]
}

func (app *application) matchmake(w http.ResponseWriter, r *http.Request) {
	body := struct {
		StaffSlot uint8 `json:"staff_slot"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		app.serverError(w, err)
		return
	}

	response := struct {
		OK     bool
		Reason string
	}{}

	_, userID := getCredentials(r)

	// Find the highest level creature on this staff
	creatures := []models.Creature{}
	err := app.db.Select(&creatures, "SELECT * FROM creature WHERE user_id = ? AND staff_slot = ? ORDER BY xp DESC", userID, body.StaffSlot)
	if err != nil || len(creatures) < 5 {
		app.serverError(w, err)
		return
	}
	highestLevel := creatures[0].GetLevel()

	// Find a match
	lobby := app.matchMaker.findOrCreateLobby(userID, highestLevel, body.StaffSlot)

	// Now loop until there's a match
	// There may be a match on the first iteration
	for start, maxSecs, it := time.Now(), 5, 0; ; it++ {
		if it&0x0f == 0 {
			// This user canceled
			if lobby.host.id == userID && lobby.host.canceled || lobby.guest.id == userID && lobby.guest.canceled {
				response.OK = false
				response.Reason = "Search canceled"
				app.writeStruct(w, response)
				return
			}

			// Opponent canceled
			if lobby.host.id == userID && lobby.guest.canceled || lobby.guest.id == userID && lobby.host.canceled {
				if lobby.host.id == userID {
					lobby.host.canceled = true
				} else {
					lobby.guest.canceled = true
				}
				app.matchMaker.cleanLobbies()
				response.OK = false
				response.Reason = "Opponent canceled"
				app.writeStruct(w, response)
				return
			}

			// Match found
			if lobby.host.id != 0 && lobby.guest.id != 0 {
				response.OK = true
				app.writeStruct(w, response)
				return
			}

			// Timeout
			if time.Since(start) > time.Duration(maxSecs)*time.Second {
				response.OK = false
				response.Reason = "Timeout"
				app.writeStruct(w, response)
				return
			}
		}
	}
}
