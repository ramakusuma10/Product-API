package models

import (
	"gorm.io/gorm"
)

type ChecklistItem struct {
	gorm.Model
	ItemName    string `json:"itemName" validate:"required"`
	Status      bool   `json:"status"`
	ChecklistID uint   `json:"checklistId"` // Foreign key to Checklist
	Checklist   Checklist `json:"-" gorm:"constraint:OnDelete:CASCADE;"` // Cascade delete
}
