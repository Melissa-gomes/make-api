package config

import "os"

type Env struct {
	APP_PORT    int
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	DB_SSLMODE  string
	DB_HOST     string
	DB_PORT     int
}

func GetEnvs() Env {
	appPort := 
	return Env{
		APP_PORT:    os.Getenv("APP_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}
}
