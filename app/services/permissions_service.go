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

type PermissionsService interface {
	GetPermissions() ([]*models.Permissions, error)
}

type permissionsService struct {
	collection *mongo.Collection
}

// Get Permissions implements SysService.
func (s *permissionsService) GetPermissions() ([]*models.Permissions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var permissions []*models.Permissions

	cursor, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		utils.Logger.Error("Error finding permissions", zap.Error(err))
		return nil, err
	}

	if err = cursor.All(ctx, &permissions); err != nil {
		utils.Logger.Error("Error decoding permissions", zap.Error(err))
		return nil, err
	}

	return permissions, nil
}

func NewPermissionsService() PermissionsService {
	return &permissionsService{
		collection: db.DB.Collection("permissions"),
	}
}
