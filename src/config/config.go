package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBConnectionString is the string that connects to MySQL
	DBConnectionString = ""

	// DBConfig is the database connection configuration
	DBConfig = "charset=utf8&parseTime=True&loc=Local"

	// Port describe where the API will be running
	Port      = 0
	SecretKey []byte
)

// Load is going to initialize ambient variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DBConnectionString = fmt.Sprintf("%s:%s@/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		DBConfig,
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
