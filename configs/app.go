package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Config struct {
	BindAddr       string         `toml:"bind_addr"`
	LogLevel       string         `toml:"log_level"`
	DatabaseConfig DatabaseConfig `toml:"database"`
}

var DB *gorm.DB

// DatabaseConfig represents db configuration
type DatabaseConfig struct {
	Host     string `toml:"DB_HOST"`
	Port     int    `toml:"DB_PORT"`
	User     string `toml:"DB_USER"`
	DBName   string `toml:"DB_NAME"`
	Password string `toml:"DB_PASSWORD"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		DatabaseConfig: DatabaseConfig{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "password",
			DBName:   "demo",
		},
	}
}

func BuildDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "myrootpassword",
		DBName:   "go-reste",
	}
}

func DbURL(DatabaseConfig DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseConfig.User,
		DatabaseConfig.Password,
		DatabaseConfig.Host,
		DatabaseConfig.Port,
		DatabaseConfig.DBName,
	)
}
