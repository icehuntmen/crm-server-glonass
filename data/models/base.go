package models

import (
	"github.com/google/uuid"
	"time"
)

type Document struct {
	ID        uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	CreateAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type DocumentCreate struct {
	ID        uuid.UUID `json:"-" bson:"_id,omitempty"`
	CreateAt  time.Time `json:"-" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt,omitempty"`
}
