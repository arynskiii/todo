package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("erro occured while running http server: %s", err.Error())
	}

}
