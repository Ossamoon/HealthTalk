package model

import "time"

type User struct {
    ID        int    `json:"id" gorm:"primaly_key"`
    Name      string `json:"name"`
    Password  string `json:"password"`
    CreatedAt time.Time
}

func CreateUser(user *User) {
	db.Create(user)
}

func FindUser(u *User) User {
	var user User
    db.Where(u).First(&user)
    return user
}