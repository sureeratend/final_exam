package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	createTb := `CREATE TABLE IF NOT EXISTS customers(
				id SERIAL PRIMARY KEY,
				name TEXT,
				email TEXT,
				status TEXT
				);`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal(err)
	}
}

// Conn xxxxx
func Conn() *sql.DB {
	return db
}

// DeleteCustomerByID xxxx
func DeleteCustomerByID(id string) error {

	stmt, err := Conn().Prepare("delete from customers where id=$1")
	if err != nil {
		return fmt.Errorf("can't prepare statement statement:%w", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("can't execiute statement statement:%w", err)
	}

	return nil
}

//GetCustomers xxxx
func GetCustomers() (*sql.Rows, error) {
	stmt, err := Conn().Prepare("select id,name,email,status from customers")
	if err != nil {
		return nil, fmt.Errorf("can't prepare statement statement:%w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("can't prepare statement statement:%w", err)
	}
	return rows, nil
}

//GetCustomerByID xxxx
func GetCustomerByID(id string) (*sql.Row, error) {

	stmt, err := Conn().Prepare("select id,name,email,status from customers where id=$1")
	if err != nil {
		return nil, fmt.Errorf("can't prepare statement statement:%w", err)
	}

	row := stmt.QueryRow(id)

	return row, nil
}

// UpdateCustomerByID xxxxxx
func UpdateCustomerByID(id string, name string, email string, status string) error {

	stmt, err := Conn().Prepare("update customers set name=$2,email=$3,status=$4 where id=$1;")
	if err != nil {
		return fmt.Errorf("can't prepare statement statement:%w", err)
	}
	_, err = stmt.Exec(id, name, email, status)
	if err != nil {
		return fmt.Errorf("can't execute statement statement:%w", err)
	}
	return nil
}
