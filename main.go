package main

import (
	"fmt"

	"github.com/KCKT0112/LiteBlog/app/config"
	"github.com/KCKT0112/LiteBlog/app/db"
	routes "github.com/KCKT0112/LiteBlog/app/routers"
	"github.com/KCKT0112/LiteBlog/app/utils"
	"go.uber.org/zap"
)

//	@title			Gin Swagger Example API
//	@version		1.0
//	@description	This is a sample server for a Swagger API with Gin.

//	@host						localhost:8083
//	@BasePath					/api
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	config.InitConfig()

	port := config.AppConfig.Server.Port
	if port == 0 {
		port = 8083 // Default port
	}

	// Initialize the logger with the configuration
	utils.InitializeLogger()
	// Connect to the MongoDB database
	db.ConnectMongoDB()

	logger := utils.Logger
	logger.Info("Starting server", zap.String("port", fmt.Sprintf("%d", port)))

	router := routes.InitRouter()

	router.Run(fmt.Sprintf(":%d", port))
}
