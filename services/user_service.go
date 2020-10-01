package services

import (
	"github.com/marcio-garcia/storecard-users-api/domain"
	"github.com/marcio-garcia/storecard-users-api/errors"
)

// CreateUser inserts a user into the database
func CreateUser(user domain.User) (*domain.User, errors.APIError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser retrives a user from storage
func GetUser(ID int64) (*domain.User, errors.APIError) {
	dao := &domain.User{ID: ID}
	if err := dao.Get(); err != nil {
		return nil, errors.CreateInternalServerError(err.Error())
	}
	return dao, nil
}
