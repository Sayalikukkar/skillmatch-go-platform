package main

import (
	"fmt"
	"log"
	"net/http"
	"taskify/db"
    "taskify/routes"
)

func main() {
    // Initialize the database connection
    db.InitDB()

    // Set up routes
    router := routes.SetupRoutes()

    // Start the HTTP server
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
