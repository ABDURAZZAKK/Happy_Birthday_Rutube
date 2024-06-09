package main

import (
	"fmt"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/dto"
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/repo"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", config.DATABASE_URL)
	if err != nil {
		panic(err)
	}
	var email, pass string
	fmt.Print("Enter email: ")

	fmt.Scanln(&email)
	fmt.Print("Enter password: ")

	fmt.Scanln(&pass)

	ur := repo.NewUserRepo(db)

	_, err = ur.Create(&dto.LoginRequest{Email: email, Password: pass})
	if err != nil {
		panic(err)
	}

	fmt.Println("Success !!!")
}
