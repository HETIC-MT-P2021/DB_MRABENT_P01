package models

import (
	"database/sql"
	"log"
)



type EmployeeByDetails struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
}
type Employees struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Email          string `json:"email"`
	OfficeCode	   string  `json:"officeCode"`
	JobTitle       string `json:"jobTitle"`
	City           string `json:"city"`
}

func (repository *Repository) GetEmployees() ([]Employees, error) {
	stmtEmp := "SELECT employeeNumber, firstName, lastName,  email, jobTitle, city FROM employees INNER JOIN offices ON employees.officeCode = offices.officeCode"

	rows, _ := repository.Conn.Query(stmtEmp)
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Email          string
		JobTitle       string
		City           string
	)

	employeesList := make([]Employees, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&EmployeeNumber, &FirstName, &LastName, &Email, &JobTitle, &City); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			employee := Employees{
				EmployeeNumber: EmployeeNumber,
				FirstName:      FirstName,
				LastName:       LastName,
				Email:          Email,
				JobTitle:       JobTitle,
				City:           City,
			}
			employeesList = append(employeesList, employee)
		default:
			return nil, err
		}
	}
	return employeesList, nil
}


	func (repository *Repository) GetEmployeeByOffice(id uint64) (EmployeeByDetails, error) {
	sqlStmt := "SELECT e.employeeNumber, e.firstName, e.lastName , e.email, e.jobTitle  FROM employees e INNER JOIN offices o ON e.officeCode = o.officeCode WHERE o.officeCode = ?;"
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Email          string
		JobTitle       string

	)
	stmt, err := repository.Conn.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&EmployeeNumber, &LastName, &FirstName, &Email, &JobTitle)
	if err != nil {
		log.Fatal(err)
	}
	customer := EmployeeByDetails{
		EmployeeNumber: EmployeeNumber,
		LastName:       LastName,
		FirstName:      FirstName,
		Email:          Email,
		JobTitle:       JobTitle,
	}
	return customer, nil
}


