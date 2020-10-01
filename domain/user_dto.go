package domain

import (
	"strings"

	"github.com/marcio-garcia/storecard-users-api/errors"
)

// User is the user model
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Password    string `json:"password"`
}

// Validate the user data
func (user *User) Validate() errors.APIError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.CreateBadRequestError("Invalid email address")
	}

	// user.Password = strings.TrimSpace(user.Password)
	// if user.Password == "" {
	// 	return errors.CreateBadRequestError("Invalid password format")
	// }
	return nil
}
