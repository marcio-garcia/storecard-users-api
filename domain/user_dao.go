package domain

import (
	"fmt"
	"log"

	"github.com/marcio-garcia/storecard-users-api/datasource/postgres/usersdb"
	"github.com/marcio-garcia/storecard-users-api/errors"
	"github.com/marcio-garcia/storecard-users-api/utils"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id;"
)

// Save - inserts a user into the database
func (user *User) Save() errors.APIError {
	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		log.Println(err.Error())
		return errors.CreateInternalServerError(err.Error())
	}
	defer statement.Close()

	dateCreated := utils.GetNow()

	//result, err := statement.Exec(user.FirstName, user.LastName, user.Email, dateCreated)
	var userID int64
	insertErr := statement.QueryRow(user.FirstName, user.LastName, user.Email, dateCreated).Scan(&userID)
	if insertErr != nil {
		return errors.CreateInternalServerError(fmt.Sprintf("Error saving the user: %s", insertErr.Error()))
	}

	user.ID = userID

	return nil
}

// Get retrieves a user from database
func (user *User) Get() error {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	return nil
}
