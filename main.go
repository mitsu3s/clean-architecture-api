package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitsu3s/clean-architecture-api/driver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
