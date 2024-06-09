package app

import (
	"os"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/middlewares"
	employee_hendlers "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/handlers"
	user_hendlers "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	SetLogger()

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

	jwt := middlewares.NewAuthMiddleware(config.SECKRET_KEY)

	uh := user_hendlers.NewUserHendlers(db)

	app.Post("/registration", uh.Create)

	app.Post("/login", uh.Login)

	app.Get("/user/get/:email", jwt, uh.GetByEmail)
	app.Get("/user/all", jwt, uh.GetAll)
	app.Delete("/user/delete/:email", jwt, uh.Delete)
	app.Post("/user/sub/:email", jwt, uh.Sub)
	app.Get("/user/subs", jwt, uh.GetAllSubs)

	eh := employee_hendlers.NewEmployeeHendlers(db)

	app.Get("/employee/get/:email", jwt, eh.GetByEmail)
	app.Get("/employee/all", jwt, eh.GetAll)
	app.Post("employee/create", jwt, eh.Create)
	app.Delete("/employee/delete/:email", jwt, eh.Delete)

	app.Listen(":3000")
}
