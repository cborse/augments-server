package models

type Actionset struct {
	SpeciesID SpeciesID `db:"species_id" json:"species_id"`
	ActionID  ActionID  `db:"action_id" json:"action_id"`
}
