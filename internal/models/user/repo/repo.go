package repo

import (
	"database/sql"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/dto"
	uerr "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/errors"
	utils "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/user/utils"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type User interface {
	GetByEmail(email string) (*dto.User, error)
	GetAll() ([]*dto.User, error)
	Create(in *dto.LoginRequest) (*dto.User, error)
	Delete(email string) (string, error)

	Sub(in *dto.SubIn) (*dto.Sub, error)
	GetAllSubs() ([]*dto.Sub, error)
}

type userRepo struct {
	*sqlx.DB
}

func NewUserRepo(conn *sqlx.DB) User {
	return &userRepo{conn}
}

func (r *userRepo) GetByEmail(email string) (*dto.User, error) {
	res := new(dto.User)
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

func (r *userRepo) GetAll() ([]*dto.User, error) {
	var res []*dto.User
	err := r.Select(&res, `SELECT email, password FROM users LIMIT $1`)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []*dto.User{}, nil
	}

	return res, nil
}

func (r *userRepo) Create(in *dto.LoginRequest) (*dto.User, error) {

	passHash, err := utils.HashPassword(in.Password)
	if err != nil {
		log.Errorf("UserService Create: %v", err)
		return nil, err
	}

	res := new(dto.User)
	err = r.Get(res, `	
	INSERT INTO users
	(email, password) 
	VALUES ($1, $2)
	RETURNING email, password`,
		in.Email, string(passHash))
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

func (r *userRepo) Sub(in *dto.SubIn) (*dto.Sub, error) {
	tx, err := r.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	res := new(dto.Sub)
	err = tx.QueryRow(`INSERT INTO subs (user, employee) VALUES ($1, $2) returning id, user, employee`, in.UserEmail, in.EmployeeEmail).
		Scan(&res.ID, &res.UserEmail, &res.EmployeeEmail)
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(`SELECT date_of_birth FROM employees WHERE email = $1;`,
		res.EmployeeEmail).Scan(&res.EmployeeDateOfBirth)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *userRepo) GetAllSubs() ([]*dto.Sub, error) {
	var res []*dto.Sub
	rows, err := r.Query(`SELECT s.id, s.user, s.employee, e.date_of_birth FROM subs AS s 
					      JOIN employees AS e ON e.email = s.employee LIMIT  $1`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sub dto.Sub
		if err := rows.Scan(&sub.ID, &sub.UserEmail, &sub.EmployeeEmail, &sub.EmployeeDateOfBirth); err != nil {
			return nil, err
		}
		res = append(res, &sub)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if res == nil {
		return []*dto.Sub{}, nil
	}

	return res, nil
}
