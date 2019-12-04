package main

import (
	"chatterbox/commands"

	"github.com/joho/godotenv"
)

func main() {
	// Loads .env file for local development
	godotenv.Load()

	// Start cobra command CLI
	commands.Execute()
}
