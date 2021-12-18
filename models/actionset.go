package models

type Actionset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	ActionID  uint32 `db:"action_id" json:"action_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
}
