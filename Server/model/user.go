package model

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name      string    `json:"name"`
    Password  string    `json:"password"`
    Email     string    `json:"email"`
    Friends   []*User   `gorm:"many2many:user_friends;"`
}

func CreateUser(user *User) {
	db.Create(user)
}

func FindUser(u *User) User {
	var user User
    db.Where(u).First(&user)
    return user
}