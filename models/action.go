package models

type Action struct {
	ID           ActionID `db:"id" json:"id"`
	Name         string   `db:"name" json:"name"`
	Type         TypeID   `db:"type" json:"type"`
	Core         bool     `db:"core" json:"core"`
	Effect       EffectID `db:"effect" json:"effect"`
	EffectChance uint8    `db:"effect_chance" json:"effect_chance"`
	Cover        CoverID  `db:"cover" json:"cover"`
	Style        StyleID  `db:"style" json:"style"`
	Detect       bool     `db:"detect" json:"detect"`
	Power        uint8    `db:"power" json:"power"`
	Speed        uint8    `db:"speed" json:"speed"`
	Energy       uint8    `db:"energy" json:"energy"`
	Accuracy     uint8    `db:"accuracy" json:"accuracy"`
	Contact      bool     `db:"contact" json:"contact"`
	Move         CoverID  `db:"move" json:"move"`
	Desc1        string   `db:"desc1" json:"desc1"`
	Desc2        string   `db:"desc2" json:"desc2"`
	Desc3        string   `db:"desc3" json:"desc3"`
}
