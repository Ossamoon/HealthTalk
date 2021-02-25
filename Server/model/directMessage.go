package model

import (
	"gorm.io/gorm"
)

type DirectMessage struct {
    gorm.Model
	FromUID   uint   `json:"from_uid"`
	ToUID     uint   `json:"to_uid"`
    Content   string `json:"content"`
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