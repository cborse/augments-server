package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"math"
	"strconv"
)

type database interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
}

type Rarity uint8

const (
	Rarity_Common = Rarity(iota)
	Rarity_Uncommon
	Rarity_Rare
)

type Type uint8

const (
	Type_None = Type(iota)
	Type_Strength
	Type_Ground
	Type_Water
	Type_Ice
	Type_Chemical
	Type_Metal
	Type_Stone
	Type_Solar
	Type_Psyche
	Type_Wind
	Type_Electric
	Type_Spirit
	Type_Fire
	Type_Illusion
)

type Action struct {
	ID           uint32 `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Type         Type   `db:"type" json:"type"`
	Core         bool   `db:"core" json:"core"`
	Effect       uint8  `db:"effect" json:"effect"`
	EffectChance uint8  `db:"effect_chance" json:"effect_chance"`
	Cover        uint8  `db:"cover" json:"cover"`
	Style        uint8  `db:"style" json:"style"`
	Detect       bool   `db:"detect" json:"detect"`
	Power        uint8  `db:"power" json:"power"`
	Speed        uint8  `db:"speed" json:"speed"`
	Energy       uint8  `db:"energy" json:"energy"`
	Accuracy     uint8  `db:"accuracy" json:"accuracy"`
	Contact      bool   `db:"contact" json:"contact"`
	Move         uint8  `db:"move" json:"move"`
	Desc1        string `db:"desc1" json:"desc1"`
	Desc2        string `db:"desc2" json:"desc2"`
	Desc3        string `db:"desc3" json:"desc3"`
}

func Action_select(db database) (*[]Action, error) {
	actions := &[]Action{}
	err := db.Select(actions, "SELECT * FROM action ORDER BY id")
	return actions, err
}

func Action_get(db database, id uint32) (*Action, error) {
	action := &Action{}
	err := db.Get(action, "SELECT * FROM action WHERE id = ?", id)
	return action, err
}

type Actionset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	ActionID  uint32 `db:"action_id" json:"action_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
}

func Actionset_select(db database) (*[]Actionset, error) {
	actionsets := &[]Actionset{}
	err := db.Select(actionsets, "SELECT * FROM actionset")
	return actionsets, err
}

func Actionset_exists(db database, speciesID uint32, actionID uint32, series uint8) (bool, error) {
	count := 0
	err := db.Get(count, "SELECT COUNT(*) FROM actionset WHERE species_id = ? AND action_id = ? AND series = ?", speciesID, actionID, series)
	return count == 1, err
}

type Creature struct {
	ID        uint64 `db:"id" json:"id"`
	UserID    uint64 `db:"user_id" json:"user_id"`
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	SeriesID  uint8  `db:"series_id" json:"series_id"`
	StaffSlot int8   `db:"staff_slot" json:"staff_slot"`
	Name      string `db:"name" json:"name"`
	Egg       bool   `db:"egg" json:"egg"`
	XP        uint32 `db:"xp" json:"xp"`
	Wins      uint32 `db:"wins" json:"wins"`
	Action1   uint32 `db:"action1" json:"action1"`
	Action2   uint32 `db:"action2" json:"action2"`
	Action3   uint32 `db:"action3" json:"action3"`
	Skill1    uint32 `db:"skill1" json:"skill1"`
	Skill2    uint32 `db:"skill2" json:"skill2"`
	Skill3    uint32 `db:"skill3" json:"skill3"`
}

func Creature_create(db database, userID uint64, species *Species, series uint8, instantHatch bool) error {
	wins := 0
	if instantHatch {
		wins = Species_getRequiredEggWins(species)
	}
	_, err := db.Exec("INSERT INTO creature (user_id, species_id, series_id, name, egg, staff_slot, wins, action1, action2) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		userID, species.ID, series, species.Name, true, -1, wins, species.Type1, species.Type2)
	return err
}

func Creature_select(db database, userID uint64) (*[]Creature, error) {
	creatures := &[]Creature{}
	err := db.Select(creatures, "SELECT * FROM creature WHERE user_id = ? ORDER BY id", userID)
	return creatures, err
}

func Creature_get(db database, id uint64) (*Creature, error) {
	creature := &Creature{}
	err := db.Get(creature, "SELECT * FROM creature WHERE id = ?", id)
	return creature, err
}

func Creature_getStorageCount(db database, userID uint64) (int, error) {
	count := 0
	err := db.Get(&count, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND egg = false", userID)
	return count, err
}

func Creature_getStaffCount(db database, userID uint64, slot uint8) (int, error) {
	count := 0
	err := db.Get(&count, "SELECT COUNT(*) FROM creature WHERE user_id = ? AND staff_slot = ?", userID, slot)
	return count, err
}

func Creature_hatch(db database, id uint64) error {
	_, err := db.Exec("UPDATE creature SET egg = false, wins = 0, staff_slot = -1 WHERE id = ?", id)
	return err
}

func Creature_learnAction(db database, creatureID uint64, actionID uint32, slot uint8) error {
	_, err := db.Exec("UPDATE creature SET action"+strconv.FormatUint(uint64(slot+1), 10)+" = ? WHERE id = ?", actionID, creatureID)
	return err
}

func Creature_learnSkill(db database, creatureID uint64, skillID uint32, slot uint8) error {
	_, err := db.Exec("UPDATE creature SET skill"+strconv.FormatUint(uint64(slot+1), 10)+" = ? WHERE id = ?", skillID, creatureID)
	return err
}

func Creature_updateStaffSlot(db database, id uint64, slot int8) error {
	_, err := db.Exec("UPDATE creature SET staff_slot = ? WHERE id = ?", slot, id)
	return err
}

type Series struct {
	ID   uint8  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func Series_select(db database) (*[]Series, error) {
	series := &[]Series{}
	err := db.Select(series, "SELECT * FROM series ORDER BY id")
	return series, err
}

type Skill struct {
	ID    uint32 `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Core  bool   `db:"core" json:"core"`
	Desc1 string `db:"desc1" json:"desc1"`
	Desc2 string `db:"desc2" json:"desc2"`
	Desc3 string `db:"desc3" json:"desc3"`
}

func Skill_select(db database) (*[]Skill, error) {
	skills := &[]Skill{}
	err := db.Select(skills, "SELECT * FROM skill ORDER BY id")
	return skills, err
}

func Skill_get(db database, id uint32) (*Skill, error) {
	skill := &Skill{}
	err := db.Get(skill, "SELECT * FROM skill WHERE id = ?", id)
	return skill, err
}

type Skillset struct {
	SpeciesID uint32 `db:"species_id" json:"species_id"`
	SkillID   uint32 `db:"skill_id" json:"skill_id"`
}

func Skillset_select(db database) (*[]Skillset, error) {
	skillsets := &[]Skillset{}
	err := db.Select(skillsets, "SELECT * FROM skillset")
	return skillsets, err
}

func Skillset_exists(db database, speciesID uint32, skillID uint32) (bool, error) {
	count := 0
	err := db.Get(count, "SELECT COUNT(*) FROM skillset WHERE species_id = ? AND skill_id", speciesID, skillID)
	return count == 1, err
}

type Species struct {
	ID           uint32 `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Rarity       Rarity `db:"rarity" json:"rarity"`
	Type1        Type   `db:"type1" json:"type1"`
	Type2        Type   `db:"type2" json:"type2"`
	Type3        Type   `db:"type3" json:"type3"`
	Height       uint32 `db:"height" json:"height"`
	Weight       uint32 `db:"weight" json:"weight"`
	InnerPower   uint8  `db:"inner_power" json:"inner_power"`
	InnerDefense uint8  `db:"inner_defense" json:"inner_defense"`
	OuterPower   uint8  `db:"outer_power" json:"outer_power"`
	OuterDefense uint8  `db:"outer_defense" json:"outer_defense"`
	MoveSpeed    uint8  `db:"move_speed" json:"move_speed"`
	ActionSpeed  uint8  `db:"action_speed" json:"action_speed"`
	Stamina      uint8  `db:"stamina" json:"stamina"`
	Accuracy     uint8  `db:"accuracy" json:"accuracy"`
	Evasion      uint8  `db:"evasion" json:"evasion"`
}

func Species_select(db database) (*[]Species, error) {
	species := &[]Species{}
	err := db.Select(species, "SELECT * FROM species ORDER BY id")
	return species, err
}

func Species_get(db database, id uint32) (*Species, error) {
	species := &Species{}
	err := db.Get(species, "SELECT * FROM species WHERE id = ?", id)
	return species, err
}

func Species_getRandom(db database, count uint) (*[]Species, error) {
	species := &[]Species{}
	err := db.Select(species, "SELECT * FROM species WHERE id != 0 ORDER BY RAND() LIMIT ?", count)
	return species, err
}

func Species_getRequiredEggWins(species *Species) int {
	return int(math.Pow(2, float64(species.Rarity)+1))
}

type Staff struct {
	UserID uint64 `db:"user_id" json:"user_id"`
	Slot   uint8  `db:"slot" json:"slot"`
	Name   string `db:"name" json:"name"`
}

const Staff_maxSlots = 6

func Staff_create(db database, userID uint64) error {
	// Find the next slot
	staffs, err := Staff_select(db, userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if staffs == nil {
		return errors.New("staffs is nil")
	}

	slot := len(*staffs)
	if len(*staffs) >= Staff_maxSlots {
		return errors.New("too many staffs")
	}

	name := "STAFF " + strconv.Itoa(slot+1)
	_, err = db.Exec("INSERT INTO staff (user_id, slot, name) VALUES (?, ?, ?)", userID, slot, name)
	return err
}

func Staff_select(db database, userID uint64) (*[]Staff, error) {
	staffs := &[]Staff{}
	err := db.Select(staffs, "SELECT * FROM staff WHERE user_id = ? ORDER BY slot", userID)
	return staffs, err
}

type User struct {
	ID           uint64 `db:"id" json:"id"`
	SteamID      uint64 `db:"steam_id" json:"steam_id"`
	Token        string `db:"token" json:"token"`
	EggCap       uint8  `db:"egg_cap" json:"egg_cap"`
	StoragePages uint8  `db:"storage_pages" json:"storage_pages"`
	Credits      uint32 `db:"credits" json:"credits"`
}

func User_create(db database, steamID uint64) error {
	_, err := db.Exec("INSERT INTO user (steam_id, egg_cap, storage_pages, credits) VALUES (?, 5, 1, 100)", steamID)
	return err
}

func User_get(db database, id uint64) (*User, error) {
	user := &User{}
	err := db.Get(user, "SELECT * FROM user WHERE id = ?", id)
	return user, err
}

func User_getBySteamID(db database, steamID uint64) (*User, error) {
	user := &User{}
	err := db.Get(user, "SELECT * FROM user WHERE steam_id = ?", steamID)
	return user, err
}

func User_assignNewToken(db database, user *User) error {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return err
	}
	user.Token = base64.URLEncoding.EncodeToString(bytes)
	_, err = db.Exec("UPDATE user SET token = ? WHERE id = ?", user.Token, user.ID)
	return err
}

type UserAction struct {
	UserID   uint64 `db:"user_id" json:"user_id"`
	ActionID uint32 `db:"action_id" json:"action_id"`
	Qty      uint8  `db:"qty" json:"qty"`
}

const userAction_maxQty = 99

func UserAction_select(db database, userID uint64) (*[]UserAction, error) {
	userActions := &[]UserAction{}
	err := db.Select(userActions, "SELECT * FROM user_action WHERE user_id = ?", userID)
	return userActions, err
}

func UserAction_add(db database, userID uint64, actionID uint32) error {
	userAction, err := userAction_get(db, userID, actionID)
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO user_action (user_id, action_id, qty) VALUES (?, ?, 1)", userID, actionID)
		return err
	} else if err != nil {
		return err
	}
	if userAction.Qty+1 > userAction_maxQty {
		return errors.New("trying to add more than allowed user actions")
	}

	_, err = db.Exec("UPDATE user_action SET qty = qty + 1 WHERE user_id = ? AND action_id = ?", userID, actionID)
	return err
}

func UserAction_remove(db database, userAction *UserAction) error {
	if userAction.Qty == 1 {
		_, err := db.Exec("DELETE FROM user_action WHERE user_id = ? AND action_id = ?", userAction.UserID, userAction.ActionID)
		return err
	}

	_, err := db.Exec("UPDATE user_action SET qty = qty - 1 WHERE user_id = ? AND action_id = ?", userAction.UserID, userAction.ActionID)
	return err
}

func userAction_get(db database, userID uint64, actionID uint32) (*UserAction, error) {
	userAction := &UserAction{}
	err := db.Get(userAction, "SELECT * FROM user_action WHERE user_id = ? AND action_id = ?", userID, actionID)
	return userAction, err
}

func UserSkill_remove(db database, userSkill *UserSkill) error {
	if userSkill.Qty == 1 {
		_, err := db.Exec("DELETE FROM user_skill WHERE user_id = ? AND skill_id = ?", userSkill.UserID, userSkill.SkillID)
		return err
	}

	_, err := db.Exec("UPDATE user_skill SET qty = qty - 1 WHERE user_id = ? AND skill_id = ?", userSkill.UserID, userSkill.SkillID)
	return err
}

type UserSkill struct {
	UserID  uint64 `db:"user_id" json:"user_id"`
	SkillID uint32 `db:"skill_id" json:"skill_id"`
	Qty     uint8  `db:"qty" json:"qty"`
}

const userSkill_maxQty = 99

func UserSkill_select(db database, userID uint64) (*[]UserSkill, error) {
	userSkills := &[]UserSkill{}
	err := db.Select(userSkills, "SELECT * FROM user_skill WHERE user_id = ?", userID)
	return userSkills, err
}

func UserSkill_add(db database, userID uint64, skillID uint32) error {
	userSkill, err := UserSkill_get(db, userID, skillID)
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO user_skill (user_id, skill_id, qty) VALUES (?, ?, 1)", userID, skillID)
		return err
	} else if err != nil {
		return err
	}
	if userSkill.Qty+1 > userSkill_maxQty {
		return errors.New("trying to add more than allowed user skills")
	}

	_, err = db.Exec("UPDATE user_skill SET qty = qty + 1 WHERE user_id = ? AND skill_id = ?", userID, skillID)
	return err
}

func UserSkill_get(db database, userID uint64, skillID uint32) (*UserSkill, error) {
	userSkill := &UserSkill{}
	err := db.Get(userSkill, "SELECT * FROM user_skill WHERE user_id = ? AND skill_id = ?", userID, skillID)
	return userSkill, err
}
