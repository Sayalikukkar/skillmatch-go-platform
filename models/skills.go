package models

type Skill struct {
	ID            int     `json:"id"`
	ProviderID    int     `json:"provider_id"` // ID of the provider offering the skill
	Category      string  `json:"category"`
	Experience    int     `json:"experience"`  // Number of years of experience
	NatureOfWork  string  `json:"nature_of_work"`
	HourlyRate    float64 `json:"hourly_rate"`
}
