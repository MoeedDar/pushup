package model

import (
	"pushup/single"
	"time"
)

type Completion struct {
	UserId      string
	Habit       string
	CompletedAt time.Time
	Amount      float32
}

func Complete(userId string, habit string, amount float32) error {
	completion := &Completion{
		UserId: userId,
		Habit:  habit,
		Amount: amount,
	}
	_, err := single.DB.Model(completion).Insert()
	return err
}

func Progress(userId string, habit string) (float32, float32, error) {
	var err error
	var user User
	if user, err = SelectUser(userId); err != nil {
		return 0, 0, err
	}
	var goal Goal
	if goal, err = SelectGoal(userId, habit); err != nil {
		return 0, 0, err
	}

	lastInterval, err := getStartTime(user.Timezone, goal.Interval)
	if err != nil {
		return 0, 0, err
	}

	var progress float32
	err = single.DB.Model(&Completion{}).
		ColumnExpr("COALESCE(SUM(amount), 0)").
		Where("user_id = ?", userId).
		Where("habit = ?", habit).
		Where("completed_at >= ?", lastInterval).
		Select(&progress)

	return progress, goal.Target, err
}

func getStartTime(timezone string, interval string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now(), err
	}

	now := time.Now().In(loc)

	switch interval {
	case "daily":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	case "weekly":
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		daysAgo := time.Duration(weekday-1) * 24 * time.Hour
		return now.Add(-daysAgo).Truncate(24 * time.Hour), nil
	case "monthly":
		return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()), nil
	default:
		return now, nil
	}
}
