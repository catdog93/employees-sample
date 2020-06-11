package entity

//emp_no      INT             NOT NULL,
//birth_date  DATE            NOT NULL,
//first_name  VARCHAR(14)     NOT NULL,
//last_name   VARCHAR(16)     NOT NULL,
//gender      ENUM ('M','F')  NOT NULL,
//hire_date   DATE            NOT NULL,

type Employee struct {
	ID        int    `db:"emp_no" json:"id"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Gender    Gender `db:"gender" json:"gender"`
	HireDate  string `db:"hire_date" json:"hireDate"`
	BirthDate string `db:"birth_date" json:"birthDate"`
}

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)
