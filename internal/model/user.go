package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) Register(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}
