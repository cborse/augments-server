package models

type Series struct {
	ID   uint8  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
