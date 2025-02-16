package appconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var CURRENT_ENV string = "test_env"

type pgxConfig struct {
	MinConns int32
	MaxConns int32
}
type dbConfig struct {
	PGXConfig *pgxConfig
	Url       string
}

type appConfig struct {
	DB *dbConfig
}

func New() *appConfig {
	if err := godotenv.Load("envs/.test_env"); err != nil {
		fmt.Print("No .env file found")
	}
	db_config := &dbConfig{
		Url: getEnv("DB_URL", ""), PGXConfig: &pgxConfig{
			MinConns: 2,
			MaxConns: 10,
		}}
	return &appConfig{
		DB: db_config,
	}

}

func getEnv(envVar string, defaultVar string) string {
	if envVar, exists := os.LookupEnv(envVar); exists {
		return envVar
	}
	fmt.Println("Value of env var was not found, use default var")
	return defaultVar
}

var Config *appConfig = New()
