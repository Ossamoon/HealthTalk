package model


import (
	"fmt"
	"errors"

	"gorm.io/gorm"
)


const (
    UNREAD uint8 = iota + 1
	PENDING
    ACCEPTED
    REFUSED
)


type FriendInvitation struct {
    gorm.Model
	FromUserID			uint	`json:"from_user_id" gorm:"index;not null;"`
	ToUserID			uint	`json:"to_user_id" gorm:"index;not null;"`
	Status				uint8	`json:"status" gorm:"index;not null;"`
}


func CreateFriendInvitation(invitation *FriendInvitation) {
    db.Create(invitation)
}


func FindFriendInvitation(inv *FriendInvitation) FriendInvitation {
    var invitation FriendInvitation
    db.Where(inv).First(&invitation)
    return invitation
}


func FindFriendInvitations(inv *FriendInvitation) []FriendInvitation {
    var invitations []FriendInvitation
    db.Where(inv).Find(&invitations)
	fmt.Println(*inv)
	fmt.Println(invitations)
    return invitations
}


func UpdateFriendInvitation(invitation *FriendInvitation, status uint8) error {
    if UNREAD <= status && status <= REFUSED {
        db.Model(&invitation).Update("status", status)
        fmt.Println("Updated FriendInvitation.Status!!")
		return nil
    } else {
		return errors.New("invalid status value to update friend invitation")
	}
}