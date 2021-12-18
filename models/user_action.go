package models

type UserAction struct {
	UserID   uint64 `db:"user_id" json:"user_id"`
	ActionID uint32 `db:"action_id" json:"action_id"`
	Qty      uint8  `db:"qty" json:"qty"`
}
