// config/mongodb.go
package db

import (
	"context"
	"time"

	"github.com/KCKT0112/LiteBlog/app/config"
	"github.com/KCKT0112/LiteBlog/app/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var DB *mongo.Database

func ConnectMongoDB() {
	host := config.AppConfig.Database.MongoDB
	db_name := config.AppConfig.Database.DB
	logger := utils.Logger

	if host == "" || db_name == "" {
		logger.Fatal("Failed to load MongoDB configuration")
		return
	}

	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Fatal("Failed to create MongoDB client", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}

	DB = client.Database(db_name)
}
