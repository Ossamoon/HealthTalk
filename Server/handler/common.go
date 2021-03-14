package handler

import (
    "time"
)

type (
	CommonCreateResponce struct {
		ID			uint
		CreatedAt   time.Time
	}

	CommonUpdateResponce struct {
		ID			uint
		UpdatedAt   time.Time
	}
)