package main

import (
	"log"

	"github.com/cristhianforerod/udemy-go-course/db"
	"github.com/cristhianforerod/udemy-go-course/handlers"
)

func main() {
	if db.CheckConnection() {
		handlers.Dispatcher()
	} else {
		log.Fatal("Error Connection!")
		return
	}
}
