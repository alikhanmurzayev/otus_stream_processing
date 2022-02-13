package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
)

type Config struct {
	Port           int    `envconfig:"PORT" default:"7775"`
	DBDriver       string `envconfig:"DB_DRIVER" default:"postgres"`
	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         int    `envconfig:"DB_PORT" default:"5432"`
	DBName         string `envconfig:"DB_NAME" default:"mydb"`
	DBUser         string `envconfig:"DB_USER" default:"myuser"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"mypassword"`
	DBSSLMode      string `envconfig:"DB_SSLMODE" default:"disable"`
	RabbitLogin    string `envconfig:"RABBIT_LOGIN" default:"guest"`
	RabbitPassword string `envconfig:"RABBIT_PASSWORD" default:"guest"`
	RabbitHost     string `envconfig:"RABBIT_HOST" default:"localhost"`
	RabbitPort     string `envconfig:"RABBIT_PORT" default:"5672"`
	QueueName      string `envconfig:"QUEUE_NAME" default:"order_events"`
}

var (
	DBConn     *gorm.DB
	RabbitConn *amqp.Connection
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
	if err := connectToRabbit(); err != nil {
		return fmt.Errorf("connectToRabbit: %w", err)
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

func connectToRabbit() (err error) {
	RabbitConn, err = amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			config.RabbitLogin,
			config.RabbitPassword,
			config.RabbitHost,
			config.RabbitPort,
		),
	)
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	return nil
}
