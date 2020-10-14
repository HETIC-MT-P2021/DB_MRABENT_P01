package models

import (
	"database/sql"
)

type Employee struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string `json:"extension"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
	City           string `json:"city"`
}
type Offices struct {
	OfficeCode  uint       `json:"officeCode"`
	City        string     `json:"city"`
	Phone       string     `json:"phone"`
	AdressLine1 string     `json:"adressLine1"`
	AdressLine2 string     `json:"adressLine2"`
	State       string     `json:"state"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postalCode"`
	Territory   string     `json:"territory"`
	Employees   []Employee `json:"employee"`
}

func (repository *Repository) GetOfficeByCode(id uint64) ([]Employee, error) {
	sqlStmt := "SELECT employeeNumber, lastName, firstName, extension, email, jobTitle FROM employees WHERE officeCode = ?"
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Extension      string
		Email          string
		JobTitle       string
	)

	rows, _ := repository.Conn.Query(sqlStmt, id)
	employeeList := make([]Employee, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&EmployeeNumber, &LastName, &FirstName, &Extension, &Email, &JobTitle); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			employee := Employee{
				EmployeeNumber: EmployeeNumber,
				LastName:       LastName,
				FirstName:      FirstName,
				Extension:      Extension,
				Email:          Email,
				JobTitle:       JobTitle,
			}
			employeeList = append(employeeList, employee)
		default:
			return nil, err
		}
	}
	return employeeList, nil
}