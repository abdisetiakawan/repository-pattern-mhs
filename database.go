package repositorypatternmhs

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection establishes a connection to the MySQL database with specific configuration settings.
// It sets the maximum number of idle connections, maximum open connections, as well as connection
// idle and lifetime durations. If the connection cannot be opened, it panics.
// Returns a pointer to the sql.DB instance representing the database connection.
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/repositorypatternmhs?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}