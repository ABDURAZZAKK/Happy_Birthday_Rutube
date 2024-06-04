package handlers

import (
	// "time"

	// "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	// "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/dto"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/repo"
	// "github.com/gofiber/fiber/v2"
	// jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type EmployeeHendlers struct {
	employeeRepo repo.Employee
}

func NewEmployeeHendlers(conn *sqlx.DB) *EmployeeHendlers {
	return &EmployeeHendlers{employeeRepo: repo.NewEmployeeRepo(conn)}
}

func (h *EmployeeHendlers) Sub() {

}
