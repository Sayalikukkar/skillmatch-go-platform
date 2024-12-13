package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
    var err error

    // Load environment variables from .env file
    err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get database credentials from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Format the DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    
    // Open the database connection
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    // Ping the database to verify the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MySQL database!")
}

func GetDB() *sql.DB {
    return db
}
