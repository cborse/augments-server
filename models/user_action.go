package models

type UserAction struct {
	UserID   uint64 `db:"user_id" json:"user_id"`
	ActionID uint32 `db:"action_id" json:"action_id"`
	Qty      uint8  `db:"qty" json:"qty"`
}

func UserAction_find(db database, userID uint64, actionID uint32) (*UserAction, error) {
	userAction := &UserAction{}
	err := db.Get(userAction, "SELECT * FROM user_action WHERE user_id = ? AND action_id = ?", userID, actionID)
	return userAction, err
}
