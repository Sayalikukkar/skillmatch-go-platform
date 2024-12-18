package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"taskify/db"
	"taskify/models"

	"github.com/gorilla/mux"
	// "log"
)

// CreateProvider creates a new provider
func CreateProvider(w http.ResponseWriter, r *http.Request) {
	var provider models.Provider
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&provider)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the provider data (this can be more complex, depending on your requirements)
	if provider.ProviderType == "" {
		http.Error(w, "Provider type is required", http.StatusBadRequest)
		return
	}

	// SQL Query to insert a new provider into the database
	query := `
		INSERT INTO providers (
			provider_type, 
			full_name, 
			email, 
			mobile_number, 
			address, 
			company_name, 
			tax_number, 
			representative
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Use db.Exec to execute the query
	_, err = db.GetDB().Exec(query, provider.ProviderType, provider.FullName, provider.Email, provider.MobileNumber, provider.Address, provider.CompanyName, provider.TaxNumber , provider.Representative)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating provider: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(provider)
}

func CreateSkill(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert into the database
	query := `INSERT INTO skills (provider_id, category, experience, nature_of_work, hourly_rate) 
	          VALUES (?, ?, ?, ?, ?)`
	result, err := db.GetDB().Exec(query, skill.ProviderID, skill.Category, skill.Experience, skill.NatureOfWork, skill.HourlyRate)
	if err != nil {
		http.Error(w, "Error inserting skill", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	skill.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(skill)
}

func UpdateSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	skillID := vars["id"]

	var skill models.Skill
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the database
	query := `UPDATE skills 
	          SET provider_id = ?, category = ?, experience = ?, nature_of_work = ?, hourly_rate = ? 
	          WHERE id = ?`
	result, err := db.GetDB().Exec(query, skill.ProviderID, skill.Category, skill.Experience, skill.NatureOfWork, skill.HourlyRate, skillID)
	if err != nil {
		http.Error(w, "Error updating skill", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No skill found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Skill updated successfully"})
}

func MakeOffer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskIDStr := vars["id"] // Extract task ID as string

	// Convert taskID to integer
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var offer models.Offer
	err = json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	offer.TaskID = taskID // Assign the integer task ID to the offer

	// Insert the offer into the database
	query := `INSERT INTO offers (task_id, provider_id, offered_rate, message) VALUES (?, ?, ?, ?)`
	_, err = db.GetDB().Exec(query, offer.TaskID, offer.ProviderID, offer.OfferedRate, offer.Message)
	if err != nil {
		http.Error(w, "Error creating offer", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Offer made successfully"})
}

func UpdateTaskProgress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var progress models.TaskProgress
	err := json.NewDecoder(r.Body).Decode(&progress)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `UPDATE task_progress SET description = ?, timestamp = ? WHERE id = ?`
	_, err = db.GetDB().Exec(query, progress.Description, progress.Timestamp, taskID)
	if err != nil {
		http.Error(w, "Error updating task progress", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task progress updated successfully"})
}

func MarkTaskCompleted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	query := `UPDATE tasks SET status = 'Completed' WHERE id = ?`
	_, err := db.GetDB().Exec(query, taskID)
	if err != nil {
		http.Error(w, "Error marking task as completed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task marked as completed"})
}
