package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"template.clevecord.me/server/constants"
	"template.clevecord.me/server/routes"
)

func main() {
	initEnv()

	r := gin.Default()
	routes.InitRoutes(r)
	var port = constants.DefaultPort
	envPort, exists := os.LookupEnv(constants.AppPortEnvKey)
	if exists {
		port = envPort
	}
	r.Run(":" + port)
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
