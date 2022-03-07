package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRole string
type UserStatue string
type UserAuth int

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

const (
	StatusActive   UserStatue = "active"
	StatusInactive UserStatue = "inactive"
)

const (
	AuthNone UserAuth = 0
	AuthKYC  UserAuth = 1
)

type User struct {
	*gorm.Model `json:"-"`
	FirstName   string     `json:"firstName,omitempty"`
	LastName    string     `json:"lastName,omitempty"`
	Email       string     `json:"email,omitempty"`
	Password    string     `json:"password,omitempty"`
	Role        UserRole   `json:"role,omitempty"`
	Status      UserStatue `json:"status,omitempty"`
	Auth        UserAuth   `json:"auth,omitempty"`
	LastLogin   time.Time  `json:"lastLogin,omitempty"`

	Balances []Balance
}

func GetUser(email string) (*User, error) {
	db := GetDB()
	var user User
	if result := db.Where("email = ?", email).Find(&user); result.RowsAffected != 1 {
		return nil, errors.New("invalid email address")
	} else {
		user.LastLogin = time.Now()
		db.Save(&user)
		return &user, nil
	}
}

func (user *User) Login() {
	db := GetDB()
	db.Model(&user).Update("last_login", time.Now())
}
