package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Permissions struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PermissionID string             `bson:"permission"`
	Description  string             `bson:"description"`
	Path         string             `bson:"path"`
	Methods      string             `bson:"methods"`
}
