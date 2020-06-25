package dao

import (
	"GoRestApi-MySql/domain"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getAllEmployees() []domain.Employee {

	db := openDbConnection()

	var employees []domain.Employee

	// Get Employees
	results, err := db.Query("SELECT * FROM employee")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var employee domain.Employee

		//For each row scan the result and convert into employee object
		err = results.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Phone)

		if err != nil {
			panic(err.Error()) // Failed to convert Database rows to Employee objects
		}
	}
	return employees
}

func openDbConnection() *sql.DB {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:bcmc1234@tcp(127.0.0.1:3306)/godemo")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}
