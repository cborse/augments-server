package models

type Creature struct {
	ID        uint64    `db:"id" json:"id"`
	UserID    uint64    `db:"user_id" json:"user_id"`
	SpeciesID SpeciesID `db:"species_id" json:"species_id"`
	StaffSlot int8      `db:"staff_slot" json:"staff_slot"`
	Name      string    `db:"name" json:"name"`
	Egg       bool      `db:"egg" json:"egg"`
	XP        uint32    `db:"xp" json:"xp"`
	Wins      uint32    `db:"wins" json:"wins"`
	Action1   ActionID  `db:"action1" json:"action1"`
	Action2   ActionID  `db:"action2" json:"action2"`
	Action3   ActionID  `db:"action3" json:"action3"`
	Skill1    SkillID   `db:"skill1" json:"skill1"`
	Skill2    SkillID   `db:"skill2" json:"skill2"`
	Skill3    SkillID   `db:"skill3" json:"skill3"`
}

func (c *Creature) CanLearnAction(actionID ActionID) bool {
	action := GetAction(actionID)
	species := GetSpecies(c.SpeciesID)
	if action.Core && (action.Type == species.Type1 || action.Type == species.Type2 || action.Type == species.Type3) {
		return true
	}
	for _, id := range species.Actionset {
		if id == actionID {
			return true
		}
	}
	return false
}

func (c *Creature) CanLearnSkill(skillID SkillID) bool {
	skill := GetSkill(skillID)
	if skill.Core {
		return true
	}
	species := GetSpecies(c.SpeciesID)
	for _, id := range species.Skillset {
		if id == skillID {
			return true
		}
	}
	return false
}
