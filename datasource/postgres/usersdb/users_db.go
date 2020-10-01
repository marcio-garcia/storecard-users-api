package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	postgresUsersUsername = "POSTGRES_USERS_USERNAME"
	postgresUsersPassword = "POSTGRES_USERS_PASSWORD"
	postgresUsersHost     = "POSTGRES_USERS_HOST"
	postgresUsersPort     = "POSTGRES_USERS_PORT"
	postgresUsersDB       = "POSTGRES_USERS_DB"
	postgresUsersSchema   = "POSTGRES_USERS_SCHEMA"
)

// Client represents the database instance
var (
	Client *sql.DB

	username = os.Getenv(postgresUsersUsername)
	password = os.Getenv(postgresUsersPassword)
	host     = os.Getenv(postgresUsersHost)
	port     = os.Getenv(postgresUsersPort)
	db       = os.Getenv(postgresUsersDB)
	schema   = os.Getenv(postgresUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, username, password, db, schema)

	var err error
	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully open")
}
