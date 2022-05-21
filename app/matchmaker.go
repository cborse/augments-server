package main

import (
	"sync"
	"time"
)

type lobbyUser struct {
	id        uint64
	level     uint32
	staffSlot uint8
}

type lobby struct {
	ready bool
	host  lobbyUser
	guest lobbyUser
}

type matchMaker struct {
	mutex   sync.Mutex
	lobbies []lobby
}

func (m *matchMaker) findOrCreateLobby(searcherID uint64, searcherLevel uint32, searcherStaffSlot uint8) *lobby {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, lobby := range m.lobbies {
		if lobby.host.level == searcherLevel {
			// Match found. Add guest data
			m.lobbies[i].ready = true
			m.lobbies[i].guest.id = searcherID
			m.lobbies[i].guest.staffSlot = searcherStaffSlot
			return &m.lobbies[i]
		}
	}

	// No match. Create a lobby
	host := lobbyUser{id: searcherID, staffSlot: searcherStaffSlot}
	m.lobbies = append(m.lobbies, lobby{host: host})

	return &m.lobbies[len(m.lobbies)-1]
}

func (m *matchMaker) waitForMatch(lobby *lobby) bool {
	for start, maxSecs, it := time.Now(), 5, 0; ; it++ {
		if it&0x0f == 0 {
			// Match found
			if lobby.ready {
				return true
			}

			// Timeout
			if time.Since(start) > time.Duration(maxSecs)*time.Second {
				return false
			}
		}
	}
}
