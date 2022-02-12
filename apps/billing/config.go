package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int    `envconfig:"PORT" default:"7777"`
	DBDriver   string `envconfig:"DB_DRIVER" default:"postgres"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"mydb"`
	DBUser     string `envconfig:"DB_USER" default:"myuser"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"mypassword"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

var (
	DBConn *gorm.DB
)

func (c Config) getDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

var config Config

func LoadConfig() error {
	err := envconfig.Process("", &config)
	if err != nil {
		return fmt.Errorf("could not envconfig.Process: %w", err)
	}
	if err := connectToDB(); err != nil {
		return fmt.Errorf("connectToDB: %w", err)
	}
	return nil
}

func connectToDB() (err error) {
	DBConn, err = gorm.Open(postgres.Open(config.getDSN()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could no connect to db: %w", err)
	}
	db, err := DBConn.DB()
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}
	return nil
}
