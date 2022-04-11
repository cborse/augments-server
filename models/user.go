package models

type User struct {
	ID           uint64 `db:"id" json:"id"`
	SteamID      uint64 `db:"steam_id" json:"steam_id"`
	Token        string `db:"token" json:"token"`
	EggCap       uint8  `db:"egg_cap" json:"egg_cap"`
	StoragePages uint8  `db:"storage_pages" json:"storage_pages"`
	Credits      uint32 `db:"credits" json:"credits"`
}

type UserAction struct {
	UserID   uint64   `db:"user_id" json:"user_id"`
	ActionID ActionID `db:"action_id" json:"action_id"`
	Qty      uint8    `db:"qty" json:"qty"`
}

type UserSkill struct {
	UserID  uint64  `db:"user_id" json:"user_id"`
	SkillID SkillID `db:"skill_id" json:"skill_id"`
	Qty     uint8   `db:"qty" json:"qty"`
}
