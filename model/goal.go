package model

import (
	"pushup/single"
	"time"
)

type Goal struct {
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Habit     string
	Interval  string
	Target    float32
}

func SelectGoal(userId string, habit string) (Goal, error) {
	goal := Goal{UserId: userId, Habit: habit}
	err := single.DB.Model(&goal).Where("user_id = ? AND habit = ?", userId, habit).Select()
	return goal, err
}

func UpsertGoal(userId string, habit string, interval string, target float32) error {
	if err := UpsertUser(userId); err != nil {
		return err
	}

	goal := &Goal{
		UserId:   userId,
		Habit:    habit,
		Interval: interval,
		Target:   target,
	}

	_, err := single.DB.Model(goal).
		OnConflict("(user_id, habit) DO UPDATE").
		Set("updated_at = now()").
		Insert()

	return err
}
