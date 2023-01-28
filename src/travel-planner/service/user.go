package service

import (
	// "fmt"
	"travel-planner/backend"
	"travel-planner/model"
	"travel-planner/util/errors"
)

func CreateUser(user *model.User) (bool, *errors.RestErr) {
	// username existed?
	success, err := backend.DB.ReadFromDB(user)
	if err != nil {
		return false, errors.NewBadRequestError("The database has error")
	}
	if !success {
		return false, errors.NewBadRequestError("The user has already exist")
	}

	// save to db

	success,err = backend.DB.SaveUser(user)
	if err != nil {
		return false, errors.NewInternalServerError("Failed to create user")
	}
	// fmt.Printf("User is added: %s\n", user.Username)
  return true, nil
}