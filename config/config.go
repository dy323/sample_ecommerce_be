package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ENVD struct {
	DB_HOST string
	SECRET_KEY string
	SAMPLE_EMAIL string
	AES_SECRET string
	HOST string
	MAIL_HOST string
	MAIL_FROM string
	MAIL_PSW string
	MAIL_PORT string
}

//load config
func EnvConfig()(*ENVD){
	godotenv.Load(".env")
	
	return &ENVD {
		DB_HOST: os.Getenv("DB_HOST"),
		SECRET_KEY: os.Getenv("SECRET_KEY"),
		SAMPLE_EMAIL: os.Getenv("SAMPLE_EMAIL"),
		AES_SECRET: os.Getenv("AES_SECRET"),
		HOST: os.Getenv("HOST"),
		MAIL_HOST: os.Getenv("MAIL_HOST"),
		MAIL_FROM: os.Getenv("MAIL_FROM"),
		MAIL_PSW: os.Getenv("MAIL_PSW"),
		MAIL_PORT: os.Getenv("MAIL_PORT"),
	}

}

