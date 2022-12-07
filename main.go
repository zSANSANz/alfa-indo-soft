package main

import (
	"os"

	service "alfa-indo-soft/service"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")
	router := service.ExtRouter(mode)
	router.Run(":" + port)
}
