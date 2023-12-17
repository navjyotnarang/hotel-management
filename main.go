package main

import (
	"log"
	"net/http"

	"github.com/gofr/gofr"
	"github.com/navjyotnarang/hotel-management/api"
)

func main() {
	router := gofr.NewRouter()
	api.SetupRoutes(router)

	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
