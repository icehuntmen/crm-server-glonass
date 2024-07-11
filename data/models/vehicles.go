package models

import (
	"github.com/google/uuid"
	"time"
)

type Vehicle struct {
	ID        uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Model     string    `json:"model,omitempty" bson:"model,omitempty"`
	Year      int       `json:"year,omitempty" bson:"year,omitempty"`
	Price     int       `json:"price,omitempty" bson:"price,omitempty"`
	Location  string    `json:"location,omitempty" bson:"location,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
