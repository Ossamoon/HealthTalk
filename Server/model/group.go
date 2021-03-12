package model


import (
    "gorm.io/gorm"
    "fmt"
)


type Group struct {
    gorm.Model
    Name      string    `json:"name" gorm:"size:50;not null;"`
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


func FindGroupWithPreload(g *Group) Group {
	var group Group
    db.Preload("Managers").Preload("Members").Where(g).First(&group)
    return group
}


func UpdateGroup(group *Group, name string) {
    if name != "" {
        db.Model(&group).Update("name", name)
        fmt.Println("Updated Group.Name!!")
    }
}