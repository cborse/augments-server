package models

import "math"

type Species struct {
	ID           uint32 `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Rarity       Rarity `db:"rarity" json:"rarity"`
	Type1        Type   `db:"type1" json:"type1"`
	Type2        Type   `db:"type2" json:"type2"`
	Height       uint32 `db:"height" json:"height"`
	Weight       uint32 `db:"weight" json:"weight"`
	InnerPower   uint8  `db:"inner_power" json:"inner_power"`
	InnerDefense uint8  `db:"inner_defense" json:"inner_defense"`
	OuterPower   uint8  `db:"outer_power" json:"outer_power"`
	OuterDefense uint8  `db:"outer_defense" json:"outer_defense"`
	MoveSpeed    uint8  `db:"move_speed" json:"move_speed"`
	ActionSpeed  uint8  `db:"action_speed" json:"action_speed"`
	Stamina      uint8  `db:"stamina" json:"stamina"`
	Accuracy     uint8  `db:"accuracy" json:"accuracy"`
	Evasion      uint8  `db:"evasion" json:"evasion"`
}

func (s *Species) GetRequiredEggWins() int {
	return int(math.Pow(2, float64(s.Rarity)+1))
}
