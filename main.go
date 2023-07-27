package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

const (
	host     = "172.17.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mbs2"
)

func dbConnect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := dbConnect()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	fetchDataFromDB(db)
}

func fetchDataFromDB(db *sql.DB) {
	// Sample query: Select data from a table
	rows, err := db.Query("SELECT data, opts FROM tasks_redisdata")
	if err != nil {
		log.Fatal("Error executing the query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var column1Value, column2Value string
		if err := rows.Scan(&column1Value, &column2Value); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Println(column1Value, column2Value)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error handling rows:", err)
	}
}
