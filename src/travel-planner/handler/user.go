package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"travel-planner/model"
	"travel-planner/service"
	"travel-planner/util/errors"

	"github.com/google/uuid"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signup request")
	w.Header().Set("Content-Type", "text/plain")

	//  Get User information from client
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		err := errors.NewBadRequestError("Cannot decode user data from client")
	  fmt.Printf("Cannot decode user data from client %v\n", err)
	}
	// if err := user.Validate(); err != nil {
	// 	return
	// }

	if user.Email == "" {
		errors.NewBadRequestError("Invalid email address")
	}
	if user.Username == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
    errors.NewBadRequestError("Invalid username")
  }
	if user.Password == "" {
    errors.NewBadRequestError("Invalid password")
  }

	user.Id = uuid.New().ID();
	


  fmt.Println(user)
	success, err := service.CreateUser(&user)
	if err != nil {
		err := errors.NewInternalServerError("Failed to save user to DB")
		fmt.Printf("Failed to save user to DB %v\n", err)
    return
	}
	if !success {
		errors.NewBadRequestError("User already exists")
		fmt.Println("User already exists")
    return
	}
	fmt.Printf("User added successfully: %s.\n", user.Username)
}