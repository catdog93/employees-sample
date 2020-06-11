package service

import (
	"employees-sample/entity"
	"employees-sample/repository"
	"errors"
)

const (
	negativeIDError = "id must be > 0"
)

func GetEmployeeByID(id int) (*entity.Employee, error) {
	if id < 1 {
		return nil, errors.New(negativeIDError)
	}
	return repository.SelectEmployeeByID(id)
}

func CreateEmployee(employee *entity.Employee) (uint64, error) {
	return repository.InsertEmployee(employee)
}
