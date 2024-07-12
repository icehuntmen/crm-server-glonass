package dto

type Role struct {
	ID          string   `json:"-" bson:"_id,omitempty"`
	Name        string   `json:"name,omitempty" example:"Admin" bson:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty" bson:"permissions,omitempty"`
}
