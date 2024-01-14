package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Note represents a note entity.
type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title" validate:"required,min=1,max=100"`
	Content    string             `json:"content" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Note_id    string             `json:"note_id"`
}

// NewNote creates a new instance of Note with the provided title and content.
func NewNote(title, content string) *Note {
	return &Note{
		ID:         primitive.NewObjectID(),
		Title:      title,
		Content:    content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Note_id:    primitive.NewObjectID().Hex(),
	}
}
