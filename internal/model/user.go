package model

import (
	"fmt"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	t "sm/pkg/token"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyCorrectPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckLogin(username string, password string, db *gorm.DB) (string, error) {
	u := User{}

	err := db.Model(User{}).Select("username", "password").Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", fmt.Errorf("this is the error 1")
	}

	err = VerifyCorrectPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf("this is the error 2")
	}

	token, err := t.GenerateToken(uint16(u.ID))

	if err != nil {
		return "", fmt.Errorf("this is the error 3")
	}

	return token, nil
}

func (u *User) Register(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {

	// Hashing password to store
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
