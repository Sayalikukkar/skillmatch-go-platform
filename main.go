package main

import (
	"fmt"
	"log"
	"net/http"
	"taskify/db"
    "taskify/routes"
    "github.com/joho/godotenv"

)

func main() {
     // Load the .env file
   
     err := godotenv.Load()
     if err != nil {
         log.Fatal("Error loading .env file")
     }

    // Initialize the database connection
    db.InitDB()

    // Set up routes
    router := routes.SetupRoutes()

    // Start the HTTP server
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
