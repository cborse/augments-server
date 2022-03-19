package models

import "math"

type Creature struct {
	ID        uint64 `db:"id" json:"id"`
	UserID    uint64 `db:"user_id" json:"user_id"`
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	StaffSlot int8   `db:"staff_slot" json:"staff_slot"`
	Name      string `db:"name" json:"name"`
	Egg       bool   `db:"egg" json:"egg"`
	XP        uint32 `db:"xp" json:"xp"`
	Wins      uint32 `db:"wins" json:"wins"`
	Action1   uint32 `db:"action1" json:"action1"`
	Action2   uint32 `db:"action2" json:"action2"`
	Action3   uint32 `db:"action3" json:"action3"`
	Skill1    uint32 `db:"skill1" json:"skill1"`
	Skill2    uint32 `db:"skill2" json:"skill2"`
	Skill3    uint32 `db:"skill3" json:"skill3"`
}

func getXPAtLevel(level int) uint32 {
	// total xp = floor(713/20736*A3^3+2*A3^2-A3)
	return uint32(float64(713)/float64(20736)*math.Pow(float64(level), 3) + 2*math.Pow(float64(level), 2) - float64(level))
}

func (c *Creature) GetLevel() int {
	for i := 2; i <= 144; i++ {
		if c.XP < getXPAtLevel(i) {
			return i - 1
		}
	}
	return 144
}
