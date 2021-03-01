package model

import (
	"gorm.io/gorm"
)

type DirectMessage struct {
    gorm.Model
	FromUserID uint   `json:"from_uid" gorm:"not null"`
	ToUserID   uint   `json:"to_uid" gorm:"not null"`
    Content    string `json:"content" gorm:"not null"`
}

type DirectMessages []DirectMessage

func CreateDirectMessage(directMessage *DirectMessage) {
    db.Create(directMessage)
}

func FindDirectMessages(dm *DirectMessage) DirectMessages {
    var directMessages DirectMessages
    db.Where(dm).Find(&directMessages)
    return directMessages
}