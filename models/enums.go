package models

type Rarity uint8

const (
	Rarity_Common = Rarity(iota)
	Rarity_Uncommon
	Rarity_Rare
)

type Type uint8

const (
	Type_None = Type(iota)
	Type_Strength
	Type_Ground
	Type_Water
	Type_Ice
	Type_Chemical
	Type_Metal
	Type_Stone
	Type_Solar
	Type_Psyche
	Type_Wind
	Type_Electric
	Type_Spirit
	Type_Fire
	Type_Illusion
)
