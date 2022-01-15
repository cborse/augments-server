package models

type Skillset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	SkillID   uint32 `db:"skill_id" json:"skill_id"`
}
