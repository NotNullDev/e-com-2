package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Surname         string
	Email           string
	AccountVerified bool
	Password        string
	Token           string
	RefreshToken    string
	AvatarUrl       string
}

type UserSession struct {
	UserID uint
	Key    string
	Value  interface{}
}
