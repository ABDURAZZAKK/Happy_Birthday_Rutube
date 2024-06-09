package repo

import (
	"database/sql"

	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/dto"
	eerr "github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/models/employee/errors"
	"github.com/jmoiron/sqlx"
)

type Employee interface {
	GetByEmail(email string) (*dto.Employee, error)
	GetAll() ([]*dto.Employee, error)
	Create(in *dto.EmployeeIn) (*dto.Employee, error)
	Delete(email string) (string, error)
}

type employeeRepo struct {
	*sqlx.DB
}

func NewEmployeeRepo(conn *sqlx.DB) Employee {
	return &employeeRepo{conn}
}

func (r *employeeRepo) GetByEmail(email string) (*dto.Employee, error) {
	res := new(dto.Employee)
	err := r.Get(res, `
				SELECT email, date_of_birth
				FROM employees 
				WHERE email = $1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, eerr.ErrEmployeeNotFound
		}
		return nil, err
	}

	return res, nil
}

func (r *employeeRepo) GetAll() ([]*dto.Employee, error) {
	var res []*dto.Employee
	err := r.Select(&res, `SELECT email, date_of_birth FROM employees LIMIT $1`)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []*dto.Employee{}, nil
	}

	return res, nil
}

func (r *employeeRepo) Create(in *dto.EmployeeIn) (*dto.Employee, error) {
	res := new(dto.Employee)
	err := r.QueryRow(`	
	INSERT INTO employees
	(email, date_of_birth) 
	VALUES ($1, $2)
	RETURNING email, date_of_birth`,
		in.Email, in.DateOfBirth).Scan(&res.Email, &res.DateOfBirth)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *employeeRepo) Delete(email string) (string, error) {
	var name string
	err := r.QueryRow("DELETE FROM employees WHERE email = $1", email).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", eerr.ErrEmployeeNotFound
		}
		return "", err
	}

	return name, nil
}
