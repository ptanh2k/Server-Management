package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	t "sm/pkg/token"
)

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyCorrectPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckLogin(username string, password string, db *gorm.DB) (string, error) {
	u := User{}

	err := db.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = VerifyCorrectPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf(err.Error())
	}

	token, err := t.GenerateToken(u.ID)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return token, nil
}

func GetUserByID(uid uint, db *gorm.DB) (User, error) {

	var u User

	if err := db.First(&u, uid).Error; err != nil {
		return u, err
	}

	u.PrepareGive()

	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
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
