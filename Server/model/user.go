package model


import (
    "gorm.io/gorm"
)


type User struct {
    gorm.Model
    Name                string    `json:"name" gorm:"index;size:50;not null;"`
    Password            string    `json:"password" gorm:"size:50;not null;"`
    Email               string    `json:"email" gorm:"uniqueIndex;size:100;not null;"`
    Friends             []*User   `gorm:"many2many:user_friends;"`
    ManagingGroups      []*Group  `gorm:"many2many:manager_groups;"`
    PerticipatingGroups []*Group  `gorm:"many2many:member_groups;"`
}


func CreateUser(user *User) {
	db.Create(user)
}


func FindUser(u *User) User {
	var user User
    db.Where(u).First(&user)
    return user
}


func FindUserWithPreload(u *User) User {
	var user User
    db.Preload("Friends").Preload("ManagingGroups").Preload("PerticipatingGroups").Where(u).First(&user)
    return user
}