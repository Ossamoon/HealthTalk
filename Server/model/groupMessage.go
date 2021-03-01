package model

import (
	"gorm.io/gorm"
)

type GroupMessage struct {
    gorm.Model
	FromUserID uint   `json:"from_uid" gorm:"not null"`
	ToGroupID   uint   `json:"to_uid" gorm:"not null"`
    Content    string `json:"content" gorm:"not null"`
}

type GroupMessages []GroupMessage

func CreateGroupMessage(groupMessage *GroupMessage) {
    db.Create(groupMessage)
}

func FindGroupMessages(gm *GroupMessage) GroupMessages {
    var groupMessages GroupMessages
    db.Where(gm).Find(&groupMessages)
    return groupMessages
}