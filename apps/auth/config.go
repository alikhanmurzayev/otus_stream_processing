package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           int    `envconfig:"PORT" default:"9999"`
	DBDriver       string `envconfig:"DB_DRIVER" default:"postgres"`
	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         int    `envconfig:"DB_PORT" default:"5432"`
	DBName         string `envconfig:"DB_NAME" default:"mydb"`
	DBUser         string `envconfig:"DB_USER" default:"myuser"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"mypassword"`
	DBSSLMode      string `envconfig:"DB_SSLMODE" default:"disable"`
	PrivateKeyPath string `envconfig:"PRIVATE_KEY_PATH" default:"config/private.pem"`
	PublicKeyPath  string `envconfig:"PUBLIC_KEY_PATH" default:"config/public.pem"`
}

var (
	DBConn *gorm.DB

	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
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
	if err := loadCryptoKeys(config.PrivateKeyPath, config.PublicKeyPath); err != nil {
		return fmt.Errorf("could not loadCryptoKeys: %w", err)
	}
	log.Println("config loaded successfully")
	return nil
}

func loadCryptoKeys(privateKeyPath, publicKeyPath string) error {
	privatePem, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("could not read file with path %s: %w", privateKeyPath, err)
	}
	publicPem, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("could not read file with path %s: %w", publicKeyPath, err)
	}
	if PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privatePem); err != nil {
		return fmt.Errorf("could not jwt.ParseRSAPrivateKeyFromPEM: %w", err)
	}
	if PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicPem); err != nil {
		return fmt.Errorf("could not jwt.ParseRSAPublicKeyFromPEM: %w", err)
	}
	return nil
}
