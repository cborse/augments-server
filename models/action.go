package models

type Action struct {
	ID           uint32 `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Type         Type   `db:"type" json:"type"`
	Core         bool   `db:"core" json:"core"`
	Effect       uint8  `db:"effect" json:"effect"`
	EffectChance uint8  `db:"effect_chance" json:"effect_chance"`
	Cover        uint8  `db:"cover" json:"cover"`
	Style        uint8  `db:"style" json:"style"`
	Detect       bool   `db:"detect" json:"detect"`
	Power        uint8  `db:"power" json:"power"`
	Speed        uint8  `db:"speed" json:"speed"`
	Energy       uint8  `db:"energy" json:"energy"`
	Accuracy     uint8  `db:"accuracy" json:"accuracy"`
	Contact      bool   `db:"contact" json:"contact"`
	Move         uint8  `db:"move" json:"move"`
	Desc1        string `db:"desc1" json:"desc1"`
	Desc2        string `db:"desc2" json:"desc2"`
	Desc3        string `db:"desc3" json:"desc3"`
}

func Action_findByID(db database, id uint32) (*Action, error) {
	action := &Action{}
	err := db.Get(action, "SELECT * FROM action WHERE id = ?", id)
	return action, err
}
