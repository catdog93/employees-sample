// Validator declares strict rules for UserBody fields. It's easy to use UserValidation() function for it.
package validator

import (
	"employees-sample/entity"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

const (
	stringFieldPattern     = "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$"
	stringFieldMinLength   = 2
	stringFieldMaxLength   = 40
	DateFormat             = "2006-01-02"
	hoursPerDay            = 24
	daysPerYear            = 365
	limitDifferenceInYears = 16
)

const (
	emptyStringError               = "error: string field can't be empty"
	lengthError                    = "error: string field has incorrect length"
	firstNameFormatError           = "error: string field hasn't a valid format"
	lastNameFormatError            = "error: string field hasn't a valid format"
	genderError                    = "error: gender field requires such values \"M\" or \"F\""
	littleDifferenceError          = "error: difference between birthDate and hireDate in years must be > %v"
	yearsDifferenceOutOfRangeError = "error: year has incorrect value"
	negativeYearsDifferenceError   = "error: hireDate must be after birthDate"
)

// Single export function for validation UserBody instance
func UserValidation(employee entity.Employee) error {
	err := firstNameValidation(employee.FirstName)
	if err != nil {
		return err
	}
	err = lastNameValidation(employee.LastName)
	if err != nil {
		return err
	}
	err = genderValidation(employee.Gender)
	if err != nil {
		return err
	}
	birthDate, err := DateValidation(DateFormat, employee.BirthDate)
	if err != nil {
		return err
	}
	hireDate, err := DateValidation(DateFormat, employee.HireDate)
	if err != nil {
		return err
	}
	years := CalculateDifferenceInYears(hireDate, birthDate)
	if years < 0 {
		return errors.New(negativeYearsDifferenceError)
	}
	if years < limitDifferenceInYears {
		return errors.New(fmt.Sprintf(littleDifferenceError, limitDifferenceInYears))
	}
	return nil
}

func firstNameValidation(firstName string) error {
	if firstName == "" {
		return errors.New(emptyStringError)
	}
	firstName = strings.TrimSpace(firstName)
	if len(firstName) < stringFieldMinLength || len(firstName) > stringFieldMaxLength {
		return errors.New(lengthError)
	}
	regex := regexp.MustCompile(stringFieldPattern)
	if !regex.MatchString(firstName) {
		return errors.New(firstNameFormatError)
	}
	return nil
}

func lastNameValidation(lastName string) error {
	if lastName == "" {
		return errors.New(emptyStringError)
	}
	lastName = strings.TrimSpace(lastName)
	if len(lastName) < stringFieldMinLength || len(lastName) > stringFieldMaxLength {
		return errors.New(lengthError)
	}
	regex := regexp.MustCompile(stringFieldPattern)
	if !regex.MatchString(lastName) {
		return errors.New(lastNameFormatError)
	}
	return nil
}

func genderValidation(gender entity.Gender) error {
	gender = entity.Gender(strings.ToUpper(string(gender)))
	if gender != entity.Male && gender != entity.Female {
		return errors.New(genderError)
	}
	return nil
}

func DateValidation(layout, dateString string) (*time.Time, error) {
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return nil, err
	}
	if date.Year() > time.Now().Year() || date.Year() < 1900 {
		return nil, errors.New(yearsDifferenceOutOfRangeError)
	}
	return &date, nil
}

func CalculateDifferenceInYears(date1, date2 *time.Time) int {
	durationDifference := date1.Sub(*date2)
	years := durationDifference.Hours() / hoursPerDay / daysPerYear
	return int(math.Ceil(years))
}
