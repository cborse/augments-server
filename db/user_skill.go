package models

type UserSkill struct {
	UserID  uint64 `db:"user_id" json:"user_id"`
	SkillID uint32 `db:"skill_id" json:"skill_id"`
	Qty     uint8  `db:"qty" json:"qty"`
}
