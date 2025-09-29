package main

import (
	"github.com/google/uuid"
	"time"
)

type MessageNonOptimized struct {
	ID            int32     // 4
	ChatSessionID uuid.UUID // 16
	WithAction    bool      // 1
	Direction     string    // 16
	FromAdmin     bool      // 1
	Status        string    // 16
	WithFile      bool      // 1
	Message       string    // 16
	IsReply       bool      // 1
	CreatedAt     time.Time // 24
}

type MessageOptimized struct {
	CreatedAt     time.Time
	Direction     string
	Status        string
	Message       string
	ID            int32
	ChatSessionID uuid.UUID
	WithAction    bool
	FromAdmin     bool
	WithFile      bool
	IsReply       bool
}
