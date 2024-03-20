package model

import (
	"pushup/single"
)

type User struct {
	Id       string
	Timezone string
}

func SelectUser(id string) (User, error) {
	user := User{Id: id}
	err := single.DB.Model(&user).Where("id = ?", id).Select()
	return user, err
}

func UpsertUser(id string) error {
	user := &User{
		Id:       id,
		Timezone: "Europe/London",
	}

	_, err := single.DB.Model(user).
		OnConflict("(id) DO UPDATE").
		Set("timezone = EXCLUDED.timezone").
		Insert()

	return err
}
