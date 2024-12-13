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

// CreateTask allows a user to create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		http.Error(w, "Error decoding task", http.StatusBadRequest)
		return
	}

	// Insert the task into the database, including currency and rate
	query := "INSERT INTO tasks (user_id, category, task_name, description, expected_start_date, expected_hours, hourly_rate, currency, currency_rate, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.GetDB().Exec(query, task.UserID, task.Category, task.TaskName, task.Description, task.ExpectedStartDate, task.ExpectedHours, task.HourlyRate, task.Currency, task.CurrencyRate, task.Status)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		http.Error(w, "Error decoding task", http.StatusBadRequest)
		return
	}

	// Update the task in the database including currency and rate
	query := "UPDATE tasks SET category = ?, task_name = ?, description = ?, expected_start_date = ?, expected_hours = ?, hourly_rate = ?, currency = ?, currency_rate = ?, status = ? WHERE id = ? AND user_id = ?"
	result, err := db.GetDB().Exec(query, task.Category, task.TaskName, task.Description, task.ExpectedStartDate, task.ExpectedHours, task.HourlyRate, task.Currency, task.CurrencyRate, task.Status, task.ID, task.UserID)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	// Check if rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error retrieving update result", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No task found to update", http.StatusNotFound)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}


