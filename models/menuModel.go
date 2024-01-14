package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Menu represents a menu entity.
type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required,min=1,max=100"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"menu_id"`
}

// NewMenu creates a new instance of Menu with the provided name, category, and dates.
func NewMenu(name, category string, startDate, endDate *time.Time) *Menu {
	return &Menu{
		ID:         primitive.NewObjectID(),
		Name:       name,
		Category:   category,
		Start_Date: startDate,
		End_Date:   endDate,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Menu_id:    primitive.NewObjectID().Hex(),
	}
}
