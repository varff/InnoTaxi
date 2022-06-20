package main

import (
	"InnoTaxi/api/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal(err)
	}
	router.POST("/login", auth.Login)
	router.POST("/sign-in", auth.Register)
	log.Fatal(router.Run(":8080"))
}
