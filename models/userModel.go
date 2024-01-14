package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Email         *string            `json:"email" validate:"required"`
	Password      *string            `json:"password" validate:"required"`
	Phone         *string            `json:"phone" validate:"required"`
	First_name    *string            `json:"first_name" validate:"required"`
	Last_name     *string            `json:"last_name" validate:"required"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
}
