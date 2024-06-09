package dto

import (
	"time"
)

type EmployeeIn struct {
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type Employee struct {
	Email       string
	DateOfBirth time.Time
}
