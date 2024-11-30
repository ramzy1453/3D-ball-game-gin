package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("MONGO_URI"))
	return os.Getenv("MONGO_URI")
}
