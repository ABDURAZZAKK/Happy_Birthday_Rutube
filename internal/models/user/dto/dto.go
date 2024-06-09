package dto

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	Email    string
	Password string
}

type Sub struct {
	ID                  int       `json:"id"`
	UserEmail           string    `json:"user_email"`
	EmployeeEmail       string    `json:"employee_email"`
	EmployeeDateOfBirth time.Time `json:"employee_date_of_birth"`
}

type SubIn struct {
	UserEmail     string `json:"user_email"`
	EmployeeEmail string `json:"employee_email"`
}
