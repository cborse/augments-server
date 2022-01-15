package models

type Staff struct {
	UserID uint64 `db:"user_id" json:"user_id"`
	Slot   uint8  `db:"slot" json:"slot"`
	Name   string `db:"name" json:"name"`
}
