package handlers

import (
    "encoding/json"
    // "fmt"
    "net/http"
    "taskify/db"
    "taskify/models"
    // "log"
)

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Set Content-Type to JSON
    w.Header().Set("Content-Type", "application/json")
    
    var user models.User

    // Parse the incoming JSON body into the User struct
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Validate required fields
    if user.FullName == "" || user.Email == "" || user.MobileNumber == "" || user.Address == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    // SQL query to insert the new user into the database
    query := `INSERT INTO users (full_name, email, mobile_number, address) VALUES (?, ?, ?, ?)`
    
    // Execute the query with the user data
    result, err := db.GetDB().Exec(query, user.FullName, user.Email, user.MobileNumber, user.Address)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    // Get the ID of the inserted user
    userID, err := result.LastInsertId()
    if err != nil {
        http.Error(w, "Error fetching user ID", http.StatusInternalServerError)
        return
    }

    // Set the ID of the user model after insertion
    user.ID = int(userID)

    // Respond with the created user
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}


