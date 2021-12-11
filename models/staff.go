package models

type Staff struct {
	ID     uint64 `db:"id" json:"id"`
	UserID uint64 `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}
