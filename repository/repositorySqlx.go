package repository

import (
	//_ "database/sql"
	"employees-sample/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type query string

const (
	getEmployeeByIDPattern query = "SELECT * FROM employees.employees WHERE employees.emp_no=%v"
)

var dataBase *sqlx.DB

// "test:test@(localhost:3306)/test"
func ConnectDB() error {
	db, err := sqlx.Open("mysql", "catdog93:knedluks@(localhost:3306)/employees")
	if err != nil {
		return err
	}
	// force a connection and test that it worked
	err = db.Ping()
	dataBase = db
	return err
}

func SelectEmployeeByID(id int) (*entity.Employee, error) {
	query := fmt.Sprintf(string(getEmployeeByIDPattern), id)
	employee := entity.Employee{}
	// fetch all places from the db, wrapped by transaction
	tx, err := dataBase.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := dataBase.Queryx(query)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&employee)
		if err != nil {
			return nil, err
		}
	}
	return &employee, nil
}

func InsertEmployee(employee *entity.Employee) (uint64, error) {
	tx := dataBase.MustBegin()
	_, err := tx.NamedExec("INSERT INTO employees.employees (first_name, last_name, gender, birth_date, hire_date) VALUES (:first_name, :last_name, :gender, :birth_date, :hire_date)", &employee)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return 0, nil
}
