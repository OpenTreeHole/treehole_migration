package config

import "os"

type MyConfig struct {
	DbUrl string
}

var Config MyConfig

func InitConfig() {
	env := os.Getenv("DB_URL")
	if env == "" {
		panic("env:DB_URL does not exist.")
	}
	Config.DbUrl = env
}
