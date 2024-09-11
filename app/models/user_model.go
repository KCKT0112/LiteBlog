package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Uid       string             `bson:"uid"`
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

type UserRefreshTokenForm struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

type UserResponse struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ProfileResponse struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
