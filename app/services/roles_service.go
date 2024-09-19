package services

import (
	"context"
	"time"

	"github.com/KCKT0112/LiteBlog/app/db"
	"github.com/KCKT0112/LiteBlog/app/models"
	"github.com/KCKT0112/LiteBlog/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type RolesService interface {
	GetRoles() ([]*models.Roles, error)
}

type rolesService struct {
	collection *mongo.Collection
}

// GetRoles implements SysService.
func (s *rolesService) GetRoles() ([]*models.Roles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var roles []*models.Roles

	cursor, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		utils.Logger.Error("Error finding roles", zap.Error(err))
		return nil, err
	}

	if err = cursor.All(ctx, &roles); err != nil {
		utils.Logger.Error("Error decoding roles", zap.Error(err))
		return nil, err
	}

	return roles, nil
}

func NewRolesService() RolesService {
	return &rolesService{
		collection: db.DB.Collection("roles"),
	}
}
