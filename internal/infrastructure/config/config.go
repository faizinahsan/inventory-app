package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Host string
	Port int
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	Level  string
	Format string
}

// Load loads configuration using Viper
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")

	// Enable environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default values
	setDefaults()

	// Read config file if exists (optional)
	if err := viper.ReadInConfig(); err != nil {
		// Config file is optional, just log the error
		fmt.Printf("No config file found, using environment variables and defaults: %v\n", err)
	}

	config := &Config{
		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
		Logger: LoggerConfig{
			Level:  viper.GetString("logger.level"),
			Format: viper.GetString("logger.format"),
		},
	}

	return config, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	// Server defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "user")
	viper.SetDefault("database.password", "myAwEsOm3pa55@w0rd")
	viper.SetDefault("database.dbname", "inventory_db")
	viper.SetDefault("database.sslmode", "disable")

	// Logger defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.format", "console")

	// Environment
	viper.SetDefault("env", "development")
}

// DatabaseURL returns the database connection URL
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// ServerAddress returns the server address
func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
