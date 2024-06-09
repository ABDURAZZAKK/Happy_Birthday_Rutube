package handlers

import (
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/dto"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func (h *UserHendlers) GetAllSubs(c *fiber.Ctx) error {
	users, err := h.userRepo.GetAllSubs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHendlers) Sub(c *fiber.Ctx) error {
	employee_email := c.Params("email")

	if employee_email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	user_email := claims["email"].(string)

	in := &dto.SubIn{UserEmail: user_email, EmployeeEmail: employee_email}

	sub, err := h.userRepo.Sub(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(sub)
}
