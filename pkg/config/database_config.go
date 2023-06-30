package config

import (
	"fuji-auth/pkg/constants"
	"os"
)

type DatabaseConfig struct {
	DbType      string
	DbHost      string
	DbUser      string
	DbPassword  string
	DbName      string
	JwtSecret   string
	JwtValidity string //In minutes
}

func LoadConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DbHost:     os.Getenv(constants.DbHost),
		DbUser:     os.Getenv(constants.DbUser),
		DbPassword: os.Getenv(constants.DbPassword),
		DbName:     os.Getenv(constants.DbName),
		DbType:     os.Getenv(constants.DbType),
	}
}
