package models

type Creature struct {
	ID        uint64 `db:"id" json:"id"`
	UserID    uint64 `db:"user_id" json:"user_id"`
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	StaffSlot int8   `db:"staff_slot" json:"staff_slot"`
	Name      string `db:"name" json:"name"`
	Egg       bool   `db:"egg" json:"egg"`
	XP        uint32 `db:"xp" json:"xp"`
	Wins      uint32 `db:"wins" json:"wins"`
}
