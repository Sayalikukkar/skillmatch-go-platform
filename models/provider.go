package models

type Provider struct {
	ID             int    `json:"id"`
	ProviderType   string `json:"provider_type"`   // Individual or Company
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	MobileNumber   string `json:"mobile_number"`
	Address        string `json:"address"`
	CompanyName    string `json:"company_name,omitempty"`  // Only if provider is a company
	TaxNumber      string `json:"tax_number,omitempty"`    // Only if provider is a company
	Representative string `json:"representative,omitempty"` // Only if provider is a company
}
