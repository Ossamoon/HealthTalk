package model

import (
	"gorm.io/gorm"
)

type GroupMessage struct {
    gorm.Model
	FromUserID uint   `json:"from_user_id" gorm:"not null"`
	ToGroupID   uint   `json:"to_group_id" gorm:"not null"`
    Content    string `json:"content" gorm:"not null"`
}

func CreateGroupMessage(groupMessage *GroupMessage) {
    db.Create(groupMessage)
}

func FindGroupMessages(gm *GroupMessage) []GroupMessage {
    var groupMessages []GroupMessage
    db.Where(gm).Find(&groupMessages)
    return groupMessages
}