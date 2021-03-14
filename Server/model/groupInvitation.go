package model


import (
	"fmt"
	"errors"

	"gorm.io/gorm"
)


type GroupInvitation struct {
    gorm.Model
	FromGroupID			uint	`json:"from_group_id" gorm:"index;not null;"`
	ToUserID			uint	`json:"to_user_id" gorm:"index;not null;"`
	Status				uint8	`json:"status" gorm:"index;not null;"`
}


func CreateGroupInvitation(invitation *GroupInvitation) {
    db.Create(invitation)
}


func FindGroupInvitation(inv *GroupInvitation) GroupInvitation {
    var invitation GroupInvitation
    db.Where(inv).First(&invitation)
    return invitation
}


func FindGroupInvitations(inv *GroupInvitation) []GroupInvitation {
    var invitations []GroupInvitation
    db.Where(inv).Find(&invitations)
    return invitations
}


func UpdateGroupInvitation(invitation *GroupInvitation, status uint8) error {
    if UNREAD <= status && status <= REFUSED {
        db.Model(&invitation).Update("status", status)
        fmt.Println("Updated GroupInvitation.Status!!")
		return nil
    } else {
		return errors.New("invalid status value to update group invitation")
	}
}