package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcio-garcia/storecard-users-api/domain"
	"github.com/marcio-garcia/storecard-users-api/errors"
	"github.com/marcio-garcia/storecard-users-api/services"
)

func getUserID(ID string) (int64, error) {
	userID, userErr := strconv.ParseInt(ID, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userID, nil
}

// CreateUser inserts a new user into the database
func CreateUser(context *gin.Context) {
	var user domain.User
	if err := context.ShouldBindJSON(&user); err != nil {
		apiError := errors.CreateBadRequestError("Invalid json body")
		context.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, &result)
}

// GetUser retrieves a user from storage
func GetUser(context *gin.Context) {
	userID, err := getUserID(context.Param("user_id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		context.JSON(http.StatusBadRequest, getErr)
		return
	}
	context.JSON(http.StatusOK, &user)
}
