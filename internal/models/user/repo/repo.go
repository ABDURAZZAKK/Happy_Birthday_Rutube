package repo

import (
	"database/sql"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/dto"
	uerr "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetByEmail(email string) (*dto.User, error)
	GetAll(limit int) ([]*dto.User, error)
	Create(in *dto.LoginRequest) (*dto.User, error)
	Delete(email string) (string, error)
}

type userRepo struct {
	*sqlx.DB
}

func NewUserRepo(conn *sqlx.DB) User {
	return &userRepo{conn}
}

func (r *userRepo) GetByEmail(email string) (*dto.User, error) {
	var res *dto.User
	err := r.Get(res, `
				SELECT email, password
				FROM users 
				WHERE email = $1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, uerr.ErrUserNotFound
		}
		return nil, err
	}

	return res, nil
}

func (r *userRepo) GetAll(limit int) ([]*dto.User, error) {
	var res []*dto.User
	err := r.Select(&res, `SELECT email, password FROM users LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []*dto.User{}, nil
	}

	return res, nil
}

func (r *userRepo) Create(in *dto.LoginRequest) (*dto.User, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("UserService Create: %v", err)
		return nil, err
	}

	var res *dto.User
	err = r.Get(res, `	
	INSERT INTO users
	(email, password) 
	VALUES ($1, $2)
	RETURNING email, password`,
		in.Email, passHash)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *userRepo) Delete(email string) (string, error) {
	var name string
	err := r.QueryRow("DELETE FROM users WHERE email = $1", email).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", uerr.ErrUserNotFound
		}
		return "", err
	}

	return name, nil
}
