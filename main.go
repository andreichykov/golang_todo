package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path // Get the path from the request URL
	if requestPath != "/todo" {
		fmt.Printf("Received request for path: %s\n", requestPath)
		io.WriteString(w, fmt.Sprintf("You requested the path: %s\n", requestPath))
		return
	}
	// TODO: @ariel check this out
	// obj := Todo{ID: 1, Title: "Sample Todo", Completed: false}
	// io.WriteString(w, obj)
}

func main() {
	testDatabase()
	http.HandleFunc("/", handler) // Register the handler for all paths
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}

func testDatabase() {
	const (
		dbUser     = "da"
		dbPassword = "" // It's better to use environment variables for passwords
		dbHost     = "localhost"
		dbPort     = "5432"
		dbName     = "testdb1"
		sslMode    = "disable"
	)

	// Connect to the default 'postgres' database to create our application DB
	postgresConnStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, sslMode)

	db, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		log.Fatalf("Failed to open default db: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping default db: %v", err)
	}

	// Check if database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if !exists {
		// Create the database if it does not exist
		_, err = db.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			log.Fatalf("Failed to create db '%s': %v", dbName, err)
		}
		log.Printf("Database '%s' created successfully.", dbName)
	} else {
		log.Printf("Database '%s' already exists.", dbName)
	}
	db.Close() // Close connection to 'postgres' db before reconnecting to the new one

	// Now, connect to the newly created (or existing) database 'testdb1'
	appDbConnStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	db, err = sql.Open("postgres", appDbConnStr)
	if err != nil {
		log.Fatalf("Failed to open app db '%s': %v", dbName, err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping app db '%s': %v", dbName, err)
	}
	log.Printf("Successfully connected to database '%s'.", dbName)

	// Create table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS example (
		id integer PRIMARY KEY,
		username varchar(255)
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table 'example': %v", err)
	}
	log.Println("Table 'example' is ready.")

	// Add a random row
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomId := r.Intn(10000)
	randomUsername := fmt.Sprintf("user_%d", randomId)

	// Use ON CONFLICT to avoid errors if the random ID already exists
	insertSQL := `INSERT INTO example (id, username) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;`
	res, err := db.Exec(insertSQL, randomId, randomUsername)
	if err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected > 0 {
		log.Printf("Inserted random row: id=%d, username=%s", randomId, randomUsername)
	} else {
		log.Printf("Row with id=%d already exists, not inserting.", randomId)
	}

	// Print table content
	log.Println("\n--- Table content: example ---")
	rows, err := db.Query("SELECT id, username FROM example ORDER BY id")
	if err != nil {
		log.Fatalf("Failed to query table: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("id: %d, username: %s\n", id, username)
		count++
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
	}
	fmt.Printf("------------------------------\nFound %d rows.\n", count)
}
