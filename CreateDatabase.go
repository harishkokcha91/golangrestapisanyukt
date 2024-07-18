package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 6500
	user     = "postgres"
	password = "admin"
	dbname   = "mynewdatabase"
)

func createDatabase() {

	// Connection string without the database name
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	// Open a connection to the PostgreSQL server
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL server!")

	// Create a new database
	createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", dbname)
	_, err = db.Exec(createDBSQL)
	if err != nil {
		log.Fatalf("Error creating database: %q", err)
	}

	fmt.Printf("Successfully created the database %s!\n", db)
}
func main() {

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Successfully connected to the database!")

	// // Example query
	// rows, err := db.Query("SELECT id, name FROM your_table")
	// if err != nil {
	// 	log.Fatalf("Error executing query: %q", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatalf("Error scanning row: %q", err)
	// 	}
	// 	fmt.Printf("ID: %d, Name: %s\n", id, name)
	// }

	// // Check for errors during row iteration
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalf("Error iterating rows: %q", err)
	// }
}
