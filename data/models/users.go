package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string    `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty" bson:"lastName,omitempty"`
}
