package config

import "os"

var DBUrl string
var Port string

func init() {
	DBUrl = os.Getenv("DB_URL")
	Port = os.Getenv("PORT")
}
