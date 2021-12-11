package main

import (
	"augments/models"
)

func (app *application) validateReplaceAction(userID uint64, creatureID uint64, actionID uint32, slot uint8) (bool, error) {
	// Valid slot
	if slot > 2 {
		return false, nil
	}

	// User ID
	creature, err := models.Creature_findByID(app.db, creatureID)
	if err != nil {
		return false, err
	}
	if creature.UserID != userID {
		return false, nil
	}

	// In inventory and qty > 0
	userAction, err := models.UserAction_find(app.db, userID, actionID)
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
	action, err := models.Action_findByID(app.db, actionID)
	if err != nil {
		return false, err
	}

	species, err := models.Species_findByID(app.db, creature.SpeciesID)
	if err != nil {
		return false, err
	}

	// Core
	if action.Core && (species.Type1 == action.Type || species.Type2 == action.Type || species.Type3 == action.Type) {
		return true, nil
	}

	// Actionset
	inActionSet, err := models.Actionset_canLearn(app.db, species.ID, actionID, creature.SeriesID)
	if inActionSet {
		return true, nil
	}

	return false, err
}
