package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Name        string
	Cost        int64
	Power       int64
	Description string
	Source      string
	Image       string
	// Tags        []string
}

type CardResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Cost        int64     `json:"cost"`
	Power       int64     `json:"power"`
	Description string    `json:"description"`
	Source      string    `json:"source"`
	Image       string    `json:"image"`
	// Tags        []string
}
