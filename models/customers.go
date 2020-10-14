package models

import (
	"database/sql"
	"log"
)

type Customer struct {
	CustomerNumber         int     `json:"customerNumber"`
	CustomerName           string  `json:"customerName"`
	ContactLastName        string  `json:"contactLastName"`
	ContactFirstName       string  `json:"contactFirstName"`
	Phone                  string  `json:"phone"`
	AddressLine1           string  `json:"addressLine1"`
	AddressLine2           string  `json:"addressLine2"`
	City                   string  `json:"city"`
	PostalCode             string  `json:"postalCode"`
	Country                string  `json:"country"`
	SalesRepEmployeeNumber int64   `json:"salesRepEmployeeNumber"`
	CreditLimit            float64 `json:"creditLimit"`
}


func (repository *Repository) GetAllCustomers() ([]Customer, error) {
	rows, _ := repository.Conn.Query("SELECT customerNumber, customerName, contactLastName, contactFirstName, phone, addressLine1, addressLine2, city, postalCode, country, salesRepEmployeeNumber, creditLimit FROM customers")
	var (
		customerNumber         int
		customerName           string
		contactLastName        sql.NullString
		contactFirstName       sql.NullString
		phone                  sql.NullString
		addressLine1           sql.NullString
		addressLine2           sql.NullString
		city                   sql.NullString
		postalCode             sql.NullString
		country                sql.NullString
		salesRepEmployeeNumber sql.NullInt64
		creditLimit            sql.NullFloat64
	)

	var customersList []Customer

	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1, &addressLine2, &city, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			customer := Customer{
				CustomerNumber:         customerNumber,
				CustomerName:           customerName,
				ContactLastName:        contactLastName.String,
				ContactFirstName:       contactFirstName.String,
				Phone:                  phone.String,
				AddressLine1:           addressLine1.String,
				AddressLine2:           addressLine2.String,
				City:                   city.String,
				PostalCode:             postalCode.String,
				Country:                country.String,
				SalesRepEmployeeNumber: salesRepEmployeeNumber.Int64,
				CreditLimit:            creditLimit.Float64,
			}
			customersList = append(customersList, customer)
		default:
			return nil, err
		}
	}
	return customersList, nil
}

func (repository *Repository) GetCustomerByNum(id int64) (*Customer, error) {

	sqlStmt := "SELECT customerNumber, customerName, contactLastName, contactFirstName, phone, addressLine1, addressLine2, city, postalCode, country, salesRepEmployeeNumber, creditLimit FROM customers WHERE customerNumber = ?"
	var (
		customerNumber         int
		customerName           string
		contactLastName        string
		contactFirstName       string
		phone                  sql.NullString
		addressLine1           sql.NullString
		addressLine2           sql.NullString
		city                   sql.NullString
		postalCode             sql.NullString
		country                sql.NullString
		salesRepEmployeeNumber sql.NullInt64
		creditLimit            sql.NullFloat64
	)
	stmt, err := repository.Conn.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1, &addressLine2, &city, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit)
	if err != nil {
		log.Fatal(err)
	}
	customer := Customer{
		CustomerNumber:         customerNumber,
		CustomerName:           customerName,
		ContactLastName:        contactLastName,
		ContactFirstName:       contactFirstName,
		Phone:                  phone.String,
		AddressLine1:           addressLine1.String,
		AddressLine2:           addressLine2.String,
		City:                   city.String,
		PostalCode:             postalCode.String,
		Country:                country.String,
		SalesRepEmployeeNumber: salesRepEmployeeNumber.Int64,
		CreditLimit:            creditLimit.Float64,
	}
	return &customer, nil
}
