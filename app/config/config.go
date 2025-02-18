package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	MongoDB string
	DB      string
}

type LoggerConfig struct {
	Level string
}

type AuthConfig struct {
	AccessTokenExpiration   int
	RefereshTokenExpiration int
	JwtSecret               string
	PasswordSalt            string
}

var AppConfig *Config

// InitConfig
func InitConfig() {
	viper.SetConfigName("config")       // Config file name
	viper.SetConfigType("yaml")         // Config file type
	viper.AddConfigPath("./app/config") // Config file path

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
