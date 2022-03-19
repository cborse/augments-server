package models

type Species struct {
	ID           SpeciesID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Rarity       RarityID  `db:"rarity" json:"rarity"`
	Type1        TypeID    `db:"type1" json:"type1"`
	Type2        TypeID    `db:"type2" json:"type2"`
	Type3        TypeID    `db:"type3" json:"type3"`
	Height       uint32    `db:"height" json:"height"`
	Weight       uint32    `db:"weight" json:"weight"`
	InnerPower   uint8     `db:"inner_power" json:"inner_power"`
	InnerDefense uint8     `db:"inner_defense" json:"inner_defense"`
	OuterPower   uint8     `db:"outer_power" json:"outer_power"`
	OuterDefense uint8     `db:"outer_defense" json:"outer_defense"`
	MoveSpeed    uint8     `db:"move_speed" json:"move_speed"`
	ActionSpeed  uint8     `db:"action_speed" json:"action_speed"`
	Stamina      uint8     `db:"stamina" json:"stamina"`
	Accuracy     uint8     `db:"accuracy" json:"accuracy"`
	Evasion      uint8     `db:"evasion" json:"evasion"`
}
