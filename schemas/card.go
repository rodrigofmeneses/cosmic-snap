package schemas

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name        string
	Cost        int64
	Power       int64
	Description string
	Source      string
	Image       string
	Tags        []string
}
