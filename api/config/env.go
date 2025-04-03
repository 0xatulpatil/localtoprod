package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("ERROR: loading env files")
	}
}
