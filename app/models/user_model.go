package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
}

type UserRegisterForm struct {
	Username string `form:"username" binding:"required,min=3,max=20"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"min=8,max=16"`
}

type UserLoginForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"min=8,max=16"`
}
