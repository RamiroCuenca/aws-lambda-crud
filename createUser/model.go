package main

// User model
type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// UserCMD model
type UserCMD struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
