package app

import (
	"os"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/middlewares"
	user_hendlers "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	SetLogger()
	// Create a new Fiber instance
	app := fiber.New()

	db, err := sqlx.Open("sqlite3", config.DATABASE_URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	sqlScheme, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(string(sqlScheme))
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key

	jwt := middlewares.NewAuthMiddleware(config.SECKRET_KEY)

	uh := user_hendlers.NewUserHendlers(db)

	// Create a Login route
	app.Post("/login", uh.Login)
	// Create a protected route
	app.Get("/protected", jwt, uh.Protected)
	// Listen on port 3000
	app.Listen(":3000")
}
