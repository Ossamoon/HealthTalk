package model

import (
    "gorm.io/gorm"
)

type Group struct {
    gorm.Model
    Name      string    `json:"name" gorm:"not null"`
	Managers  []*User   `gorm:"many2many:manager_groups;"`
	Members   []*User   `gorm:"many2many:member_groups;"`
}

func CreateGroup(group *Group) {
	db.Create(group)
}

func FindGroup(g *Group) Group {
	var group Group
    db.Where(g).First(&group)
    return group
}