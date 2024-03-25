package initializers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	// Load the .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}
