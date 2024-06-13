package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	log.Println(env, "envenvenvenvenvenvenvenvenvenvenvenv")

	if env != "production" {
		var envPath string
		switch env {
		case "test":
			projectRoot, err := filepath.Abs("../")
			if err != nil {
				log.Fatal("Error determining the project root directory: ", err)
			}
			envPath = filepath.Join(projectRoot, ".env")
			log.Println("Project root directory:", projectRoot)

		default:
			envPath = ".env"
		}

		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}
}
