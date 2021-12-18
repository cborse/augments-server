package models

type Creature struct {
	ID        uint64 `db:"id" json:"id"`
	UserID    uint64 `db:"user_id" json:"user_id"`
	StaffID   uint64 `db:"staff_id" json:"staff_id"`
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
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
