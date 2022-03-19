package models

type Skillset struct {
	SpeciesID SpeciesID `db:"species_id" json:"species_id"`
	SkillID   SkillID   `db:"skill_id" json:"skill_id"`
}
