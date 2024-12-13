package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taskify/db"
	"taskify/models"
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
