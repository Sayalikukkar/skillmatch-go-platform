package models

type Offer struct {
	ID          int     `json:"id"`
	TaskID      int     `json:"task_id"`
	ProviderID  int     `json:"provider_id"`
	OfferedRate float64 `json:"offered_rate"`
	Message     string  `json:"message"`
	Status      string  `json:"status"` // e.g., "Pending", "Accepted", "Rejected"
	CreatedAt   string  `json:"created_at"`
}
