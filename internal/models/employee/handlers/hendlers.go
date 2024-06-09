package handlers

import (
	"time"

	"cloud.google.com/go/civil"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/dto"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type EmployeeHendlers struct {
	employeeRepo repo.Employee
}

func NewEmployeeHendlers(conn *sqlx.DB) *EmployeeHendlers {
	return &EmployeeHendlers{employeeRepo: repo.NewEmployeeRepo(conn)}
}

func (h *EmployeeHendlers) Create(c *fiber.Ctx) error {
	type req struct {
		Email       string     `json:"email"`
		DateOfBirth civil.Date `json:"date_of_birth"`
	}
	in := new(req)
	if err := c.BodyParser((in)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	t := time.Date(in.DateOfBirth.Year, in.DateOfBirth.Month, in.DateOfBirth.Day, 0, 0, 0, 0, time.Local)
	employee, err := h.employeeRepo.Create(&dto.EmployeeIn{Email: in.Email, DateOfBirth: t})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(employee)
}

func (h *EmployeeHendlers) GetByEmail(c *fiber.Ctx) error {
	in := c.Params("email")
	if in == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	employee, err := h.employeeRepo.GetByEmail(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(employee)
}

func (h *EmployeeHendlers) GetAll(c *fiber.Ctx) error {
	employees, err := h.employeeRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(employees)
}

func (h *EmployeeHendlers) Delete(c *fiber.Ctx) error {
	in := c.Params("email")
	if in == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	employee, err := h.employeeRepo.Delete(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(employee)
}
