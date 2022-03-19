package models

type Skill struct {
	ID    SkillID `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Core  bool    `db:"core" json:"core"`
	Desc1 string  `db:"desc1" json:"desc1"`
	Desc2 string  `db:"desc2" json:"desc2"`
	Desc3 string  `db:"desc3" json:"desc3"`
}
