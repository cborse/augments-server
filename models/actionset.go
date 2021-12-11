package models

import "database/sql"

type Actionset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	ActionID  uint32 `db:"action_id" json:"action_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
}

func Actionset_canLearn(db database, speciesID uint32, actionID uint32, seriesID uint8) (bool, error) {
	actionSet := &Actionset{}
	err := db.Get(actionSet, "SELECT * FROM actionset WHERE species_id = ? AND action_id = ? AND series_id = ?", speciesID, actionID, seriesID)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
