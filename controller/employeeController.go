package controller

import (
	"employees-sample/entity"
	"employees-sample/service"
	"employees-sample/validator"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"regexp"
	"strconv"
)

var (
	EmployeesSlashURI string = "/employees/"
	EmployeesURI      string = "/employees"

	EmployeesIDPattern = "^/employees/[0-9]*$"
)

const (
	employeeWithIDNotFoundError = "erorr: employee with id = %v doesn't exist"
)

type ReplaceIdRequestBody struct {
	*entity.Employee `json:"employee"`
}

func GetEmployeeByID(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		result, err := checkEmployeeID(r.RequestURI)
		if err != nil {
			fmt.Fprintf(rw, "%s", err)
			return
		}
		if result {
			stringURL := r.URL.String()
			id, err := strconv.Atoi(path.Base(stringURL))
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
				return
			}
			employee, err := service.GetEmployeeByID(id)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
				return
			}
			if employee.ID == 0 { // employee with such id doesn't exist
				rw.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(rw, "%s", fmt.Sprintf(employeeWithIDNotFoundError, id))
				return
			}
			bytes, err := json.Marshal(&employee)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
				return
			}
			fmt.Fprintf(rw, "%s", bytes)
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	}
	if r.Method == http.MethodPost {
		CreateEmployee(rw, r)
	}
	if r.Method == http.MethodPut {

	}
	if r.Method == http.MethodDelete {

	}
}

func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.URL.Path == EmployeesSlashURI || r.URL.Path == EmployeesURI {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
			}
			defer r.Body.Close()
			inputEmployee := entity.Employee{}
			err = json.Unmarshal(body, &inputEmployee)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
			}
			err = validator.UserValidation(inputEmployee)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
				return
			}
			id, err := service.CreateEmployee(&inputEmployee)
			if err != nil {
				fmt.Fprintf(rw, "%s", err)
			}
			fmt.Fprintf(rw, "%s", strconv.Itoa(int(id)))
			rw.WriteHeader(http.StatusCreated)
		}
	}
}

func checkEmployeeID(uri string) (bool, error) {
	return regexp.MatchString(EmployeesIDPattern, uri)
}
