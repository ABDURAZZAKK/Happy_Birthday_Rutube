package handlers

import (
	"time"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/dto"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/repo"
	utils "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/utils"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type UserHendlers struct {
	userRepo repo.User
}

func NewUserHendlers(conn *sqlx.DB) *UserHendlers {
	return &UserHendlers{userRepo: repo.NewUserRepo(conn)}
}

func (h *UserHendlers) GetByEmail(c *fiber.Ctx) error {
	in := c.Params("email")
	if in == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	user, err := h.userRepo.GetByEmail(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *UserHendlers) GetAll(c *fiber.Ctx) error {
	users, err := h.userRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHendlers) Create(c *fiber.Ctx) error {
	in := new(dto.LoginRequest)
	if err := c.BodyParser((in)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.userRepo.Create(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

// Login route
func (h *UserHendlers) Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(dto.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials
	user, err := h.userRepo.GetByEmail(loginRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = utils.ValidatePassword(user.Password, loginRequest.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24
	// Create the JWT claims, which includes the user
	claims := jtoken.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SECKRET_KEY))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(dto.LoginResponse{
		Token: t,
	})
}

// Protected route
func (h *UserHendlers) Delete(c *fiber.Ctx) error {
	in := c.Params("email")
	if in == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	user, err := h.userRepo.Delete(in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
