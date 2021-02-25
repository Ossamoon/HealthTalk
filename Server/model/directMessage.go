package model

import (
	"time"
)

type DirectMessage struct {
	ID        int    `json:"id" gorm:"praimaly_key"`
    FromUID   int    `json:"uid_from"`
	ToUID     int    `json:"uid_to"`
    Content   string `json:"content"`
	CreatedAt time.Time
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