package main

import (
	"debts/internal /handler"
	"debts/internal /repository"
	"debts/internal /service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := repository.NewClient(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		Password: "qwerty",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Can't connect to db. Error: %s", err.Error())
	}
	repo := repository.NewRepo(db)
	serv := service.NewService(repo)
	handl := handler.NewHandler(serv)

	router := gin.New()

	handl.InitRoute(router)

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatalf("The server finish with err: %s", err.Error())
	}
	log.Println("The server shut down")
}
