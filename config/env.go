package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(".env.local"); err != nil {
			fmt.Println("ERROR: loading env files", err)
		}
	}
}
