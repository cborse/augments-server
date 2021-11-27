package models

type Skillset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	SkillID   uint32 `db:"skill_id" json:"skill_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
}
