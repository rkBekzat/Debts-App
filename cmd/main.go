package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.New()

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatalf("The server finish with err: %s", err.Error())
	}
	log.Println("The server shut down")
}
