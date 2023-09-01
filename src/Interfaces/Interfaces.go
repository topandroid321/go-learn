package Interfaces

type SystemInterface struct {
	CorsAllowOrigin string
	CorsAllowHeader string
	Port            int
}

type AddCustomer struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
