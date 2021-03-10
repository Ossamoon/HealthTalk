package model

import (
	"gorm.io/gorm"
)

type DirectMessage struct {
    gorm.Model
	FromUserID uint   `json:"from_user_id" gorm:"index;not null;"`
	ToUserID   uint   `json:"to_user_id" gorm:"index;not null;"`
    Content    string `json:"content" gorm:"not null"`
}

func CreateDirectMessage(directMessage *DirectMessage) {
    db.Create(directMessage)
}

func FindDirectMessages(dm *DirectMessage) []DirectMessage {
    var directMessages []DirectMessage
    db.Where(dm).Find(&directMessages)
    return directMessages
}