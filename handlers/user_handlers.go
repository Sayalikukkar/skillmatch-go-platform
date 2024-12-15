package handlers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"taskify/db"
	"taskify/models"
	"github.com/gorilla/mux"
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
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (user_id, category, task_name, description, expected_start_date, expected_hours, hourly_rate, currency, currency_rate, status) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.GetDB().Exec(query, task.UserID, task.Category, task.TaskName, task.Description, task.ExpectedStartDate, task.ExpectedHours, task.HourlyRate, task.Currency, task.CurrencyRate, task.Status)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	task.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the URL
	vars := mux.Vars(r)
	taskID := vars["id"]

	// Decode the request body into the Task struct
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure UserID is provided in the payload
	if task.UserID == 0 {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Update the task in the database
	query := `
		UPDATE tasks
		SET category = ?, task_name = ?, description = ?, expected_start_date = ?, 
		    expected_hours = ?, hourly_rate = ?, currency = ?, currency_rate = ?, status = ?
		WHERE id = ? AND user_id = ?
	`
	result, err := db.GetDB().Exec(query, task.Category, task.TaskName, task.Description, task.ExpectedStartDate,
		task.ExpectedHours, task.HourlyRate, task.Currency, task.CurrencyRate, task.Status, taskID, task.UserID)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	// Check if the update was applied
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No task found with the given ID and User ID", http.StatusNotFound)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func AcceptOffer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offerID := vars["id"]

	query := `UPDATE offers SET status = 'Accepted' WHERE id = ?`
	_, err := db.GetDB().Exec(query, offerID)
	if err != nil {
		http.Error(w, "Error accepting offer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Offer accepted successfully"})
}

func RejectOffer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offerID := vars["id"]

	query := `UPDATE offers SET status = 'Rejected' WHERE id = ?`
	_, err := db.GetDB().Exec(query, offerID)
	if err != nil {
		http.Error(w, "Error rejecting offer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Offer rejected successfully"})
}

func AcceptTaskCompletion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	query := `UPDATE tasks SET status = 'Approved' WHERE id = ?`
	_, err := db.GetDB().Exec(query, taskID)
	if err != nil {
		http.Error(w, "Error accepting task completion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task completion accepted"})
}



