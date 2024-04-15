// Package config contains the configuration for the application.
package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"xrChat_backend/internal/model"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// Config config struct.
type Config struct {
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	EmCode    string `yaml:"emCode"`
	FromEmail string `yaml:"fromEmail"`
	Mysql     struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	}
	Logging struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func GetConfig() *Config {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("config file not found")
	}

	var config = &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic("failed to unmarshal config file:" + err.Error())
	}

	return config
}

var DB *gorm.DB
var RedisClient *redis.Client
var FromEmail string
var EmCode string
var Port int

// Global init some global config.
func Global() {
	config := GetConfig()
	output := os.Stdout
	// production mode.
	if config.Mode == "prod" {
		var err error
		output, err = os.OpenFile(config.Logging.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Failed to open log file: %v" + err.Error())
		}
	}
	// Setup logging.
	// Example: logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewTextHandler(output, nil))
	slog.SetDefault(logger)
	// Connect to mysql.
	// Example: db, err := gorm.Open("username:password@tcp(host:port)/database",&gorm.Config{})
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Db)), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to mysql", "error", err)
		panic(err)
	}
	slog.Info("Connected to mysql")
	// Migrate the database.
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to auto migrate user")
	}
	slog.Info("Migrated database")
	// Initialize the redis client.
	// Example: redisClient := redis.NewClient(&redis.Options{
	// Addr:     "localhost:6379",})
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.Redis.Host + ":" + fmt.Sprint(config.Redis.Port),
	})
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to ping redis")
	}
	slog.Info("Connected to redis")
	FromEmail = config.FromEmail
	EmCode = config.EmCode
	Port = config.Port
}
