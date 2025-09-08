package config

import "os"

var (
	JWT_SECRET           string
	REFRESH_TOKEN_SECRET string
)

func LoadConfig() {
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		Logger.Printf("JWT_SECRET not set in environment")
	}
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")
	if REFRESH_TOKEN_SECRET == "" {
		Logger.Printf("REFRESH_TOKEN_SECRET not set in environment")
	}
}
