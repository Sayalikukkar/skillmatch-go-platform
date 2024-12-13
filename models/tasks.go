package models

import "time"

type Task struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"` // ID of the user who created the task
	ProviderID       int       `json:"provider_id,omitempty"` // ID of the provider assigned (can be null initially)
	Category         string    `json:"category"`
	TaskName         string    `json:"task_name"`
	Description      string    `json:"description"`
	ExpectedStartDate time.Time `json:"expected_start_date"`
	ExpectedHours    int       `json:"expected_hours"`
	HourlyRate       float64   `json:"hourly_rate"`
	Currency         string    `json:"currency"`          // Currency type (e.g., USD, INR)
	CurrencyRate     float64   `json:"currency_rate"`     // Conversion rate for the selected currency
	Status           string    `json:"status"`            // e.g., "created", "in_progress", "completed"
}
