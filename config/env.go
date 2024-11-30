package config

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGO_URI")
}
