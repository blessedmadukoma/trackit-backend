package util

import (
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application
type Config struct {
	GinMode              string        `mapstructure:"GIN_MODE"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or env variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error loading .env:", err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// LoadEnvConfig reads configuration from file or env variables
func LoadEnvConfig(path string) (config Config) {
	config.GinMode = os.Getenv("GIN_MODE")
	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	config.AccessTokenDuration, _ = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	config.RefreshTokenDuration, _ = time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))

	return config
}
