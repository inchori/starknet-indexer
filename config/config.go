package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	StarknetMainnetRpcUrl string
	StarknetSepoliaRpcUrl string
	StarknetRpcApiKey     string
	DbHost                string
	DbPort                string
	DbName                string
	DbUser                string
	DbPassword            string
	Port                  string
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		StarknetMainnetRpcUrl: os.Getenv("STARKNET_MAINNET_RPC_URL"),
		StarknetSepoliaRpcUrl: os.Getenv("STARKNET_SEPOLIA_RPC_URL"),
		StarknetRpcApiKey:     os.Getenv("STARKNET_RPC_API_KEY"),
		DbHost:                os.Getenv("DB_HOST"),
		DbPort:                os.Getenv("DB_PORT"),
		DbName:                os.Getenv("DB_NAME"),
		DbUser:                os.Getenv("DB_USER"),
		DbPassword:            os.Getenv("DB_PASSWORD"),
		Port:                  os.Getenv("PORT"),
	}
}
