package model

import "time"

type Talk struct {
	ID        int    `json:"id" gorm:"praimaly_key"`
    UID_from  int    `json:"uid_from"`
	UID_to    int    `json:"uid_to"`
    Content   string `json:"content"`
	CreatedAt time.Time
}

type Talks []Talk