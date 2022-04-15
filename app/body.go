package main

import "augments/models"

type LoginRequest struct {
	SteamID     uint64 `json:"steam_id"`
	SteamTicket string `json:"steam_ticket"`
}

type LoginResponse struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type GetDataResponse struct {
	User        *models.User         `json:"user"`
	Staffs      *[]models.Staff      `json:"staffs"`
	Creatures   *[]models.Creature   `json:"creatures"`
	UserActions *[]models.UserAction `json:"user_actions"`
	UserSkills  *[]models.UserSkill  `json:"user_skills"`
}

type AssignRequest struct {
	CreatureID uint64 `json:"creature_id"`
	StaffSlot  uint8  `json:"staff_slot"`
}

type UnassignRequest struct {
	CreatureID uint64 `json:"creature_id"`
}

type HatchEggRequest struct {
	CreatureID uint64 `json:"creature_id"`
}

type LearnActionRequest struct {
	CreatureID uint64          `json:"creature_id"`
	ActionID   models.ActionID `json:"action_id"`
	Slot       uint8           `json:"slot"`
}

type LearnSkillRequest struct {
	CreatureID uint64         `json:"creature_id"`
	SkillID    models.SkillID `json:"skill_id"`
	Slot       uint8          `json:"slot"`
}
