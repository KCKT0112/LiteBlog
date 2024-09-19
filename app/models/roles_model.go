package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Roles struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RoleID      string             `bson:"primitive_id"`
	Permissions string             `bson:"permissions"`
}
