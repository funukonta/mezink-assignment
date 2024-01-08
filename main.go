package main

import (
	"log"
)

func main() {
	repo, err := NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	// initiate db
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	// Create Server running on port :8080
	api := NewAPIServer(":8080", repo)

	log.Println("Server running on port", api.ListenPort)
	err = api.Run() // Run API Server
	if err != nil {
		log.Println("Server can't run :", err.Error())
		return
	}
}
