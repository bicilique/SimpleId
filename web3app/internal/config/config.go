package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	ClientPathNode1 string
	WalletNode1     string
	ClientPathNode2 string
	WalletNode2     string
	MKey            string
	MNonce          string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	var envFile string

	switch env {
	case "production":
		envFile = ".env.production"
	default:
		envFile = ".env.development"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}

	return &Config{
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		ClientPathNode1: os.Getenv("IPC_PATH_NODE1"),
		WalletNode1:     os.Getenv("WALLET_NODE1"),
		ClientPathNode2: os.Getenv("IPC_PATH_NODE2"),
		WalletNode2:     os.Getenv("WALLET_NODE2"),
		MKey:            os.Getenv("MASTER_KEY"),
		MNonce:          os.Getenv("MASTER_NONCE"),
	}
}
