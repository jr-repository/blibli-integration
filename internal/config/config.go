package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL      string
	APIRouter    string
	SellerKey    string
	ClientID     string
	ClientSecret string
	StoreID      string
	Port         string
}

// CONFIGURATION INITIALIZATION
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		BaseURL:      os.Getenv("BLIBLI_BASE_URL"),
		APIRouter:    os.Getenv("BLIBLI_API_ROUTER"),
		SellerKey:    os.Getenv("BLIBLI_SELLER_KEY"),
		ClientID:     os.Getenv("BLIBLI_CLIENT_ID"),
		ClientSecret: os.Getenv("BLIBLI_CLIENT_SECRET"),
		StoreID:      os.Getenv("BLIBLI_STORE_ID"),
		Port:         os.Getenv("PORT"),
	}

	if cfg.BaseURL == "" || cfg.SellerKey == "" || cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, errors.New("missing required blibli environment variables")
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg, nil
}