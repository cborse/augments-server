package main

import (
	"augments/models"
	"database/sql"
)

func (app *application) validateReplaceAction(userID uint64, creatureID uint64, actionID uint32, slot uint8) (bool, error) {
	// Valid slot
	if slot > 2 {
		return false, nil
	}

	// User ID
	creature := &models.Creature{}
	err := app.db.Get(creature, "SELECT * FROM creature WHERE id = ?", creatureID)
	if err != nil {
		return false, err
	}
	if creature.UserID != userID {
		return false, nil
	}

	// In inventory and qty > 0
	userAction := &models.UserAction{}
	err = app.db.Get(userAction, "SELECT * FROM user_action WHERE user_id = ? AND action_id = ?", userID, actionID)
	if err != nil {
		return false, err
	} else if userAction.Qty == 0 {
		return false, nil
	}

	// Not egg
	if creature.Egg {
		return false, nil
	}

	// Creature doesn't already know this action
	if creature.Action1 == actionID || creature.Action2 == actionID || creature.Action3 == actionID {
		return false, nil
	}

	// Species can learn this action
	action := &models.Action{}
	err = app.db.Get(action, "SELECT * FROM action WHERE id = ?", actionID)
	if err != nil {
		return false, err
	}

	species := &models.Species{}
	err = app.db.Get(species, "SELECT * FROM species WHERE id = ?", creature.SpeciesID)
	if err != nil {
		return false, err
	}

	// Core
	if action.Core && (species.Type1 == action.Type || species.Type2 == action.Type || species.Type3 == action.Type) {
		return true, nil
	}

	// Actionset
	actionset := &models.Actionset{}
	err = app.db.Get(actionset, "SELECT * FROM actionset WHERE species_id = ? AND action_id = ? AND series_id = ?", creature.SpeciesID, actionID, creature.SeriesID)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
