package main

import (
	"log"
	"test_project/internal/api"
	"test_project/internal/repository"
	"test_project/internal/utils"
	"test_project/pkg"
)

func main() {
	// Load Env Variables
	envList, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading env variables: %v", err)
	}
	//Preparing DSP for response testing
	db, err := pkg.InitDspDatabase(envList.PostgresUser,
		envList.PostgresPass, envList.PostgresDBName, envList.PostgresHost, envList.PostgresPort)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()
	log.Println("Database initialized successfully")
	// Add test ads into database
	err = repository.AddTestAds(db)
	if err != nil {
		log.Fatalf("Error adding test ads into dsp database: %v", err)
	}
	// Start Server
	err = api.StartServer(envList.ServerHost, envList.ServerPort, db)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
