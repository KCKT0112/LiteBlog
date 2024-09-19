package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Roles struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RoleID      string             `bson:"role_id"`
	Permissions []string           `bson:"permissions"`
}
