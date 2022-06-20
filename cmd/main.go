package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"InnoTaxi/api/auth"
	"InnoTaxi/api/server"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatalf("Failed to load environment: %s", err)
	}
	router.POST("/login", auth.Login)
	router.POST("/sign-in", auth.Register)
	router.GET("/rating", server.CheckRate)
	log.Fatal(router.Run(":8080"))
}
