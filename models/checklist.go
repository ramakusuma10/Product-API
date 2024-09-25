package models

import (
	"gorm.io/gorm"
)

type Checklist struct {
	gorm.Model
	Name           string          `json:"name" validate:"required"`
	ChecklistItems []ChecklistItem `json:"items" gorm:"foreignKey:ChecklistID"`
}
