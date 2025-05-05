package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AllEnvs struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	JWTSecrect        string
	JWTExpirationTime int
}

var Env *AllEnvs = configInit()

func configInit() *AllEnvs {

	err := godotenv.Load()
	if err != nil {

		log.Fatal("Problem in loading .env files")
	}

	return &AllEnvs{
		DbHost:     getEnv("DB_HOST"),
		DbPort:     getEnv("DB_PORT"),
		DbUser:     getEnv("DB_USER"),
		DbPassword: getEnv("DB_PASSWORD"),
		DbName:     getEnv("DB_NAME"),

		JWTSecrect:        getEnv("JWT_SECRECT"),
		JWTExpirationTime: getEnvInt("JWT_EXPIRATION_TIME"),
	}
}

func getEnv(envName string) string {
	value, ok := os.LookupEnv(envName)
	if !ok || value == "" {
		log.Fatal("env missing")
	}
	return value
}

func getEnvInt(envName string) int {
	value, ok := os.LookupEnv(envName)
	if !ok || value == "" {
		log.Fatal("env missing")
	}
	num, _ := strconv.Atoi(value)
	return num
}
