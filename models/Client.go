package models

type Client struct {
	ClientId  int    `json:"client_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
