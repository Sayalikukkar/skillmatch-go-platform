package models

type User struct {
	ID            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	MobileNumber  string `json:"mobile_number"`
	Address       string `json:"address"`
}
