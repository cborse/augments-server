package models

type User struct {
	ID           uint64 `db:"id" json:"id"`
	SteamID      uint64 `db:"steam_id" json:"steam_id"`
	Token        string `db:"token" json:"token"`
	EggCap       uint8  `db:"egg_cap" json:"egg_cap"`
	StoragePages uint8  `db:"storage_pages" json:"storage_pages"`
}
