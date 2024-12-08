package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Envs = initConfig()

// initConfig reads the .env file and loads the environment variables into the Config struct.
func initConfig() Config {

	// Adjust path based on the current directory
	envFilePath := "./.env"

	// Open the .env file
	file, err := os.Open(envFilePath)
	if err != nil {
		fmt.Println("Error opening .env file:", err)
		return Config{}
	}
	defer file.Close()

	// Initialize a new Config struct
	config := Config{}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value based on '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		// Trim spaces
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Assign values to the Config struct
		switch key {
		case "DBHost":
			config.DBHost = value
		case "DBPort":
			config.DBPort = value
		case "DBUser":
			config.DBUser = value
		case "DBPassword":
			config.DBPassword = value
		case "DBName":
			config.DBName = value
		}
	}

	// Handle scanner error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .env file:", err)
	}

	return config
}
