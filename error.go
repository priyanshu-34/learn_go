package main

import (
	"errors"
	"fmt"
)

var UserNotFoundError error = errors.New("User not found")

func getUser(id string) ( string, error){
	if(id == ""){
		return "", UserNotFoundError
	}
	return "User: " + id, nil
}

type customError struct {
	Field string
	Message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Invalid %s: %q", e.Field, e.Message)
}

func validateUserInput(input string) error {
	if input == "" {
		return &customError{
			Field:   "input",
			Message: "Input cannot be empty",
		}
	}
	return nil
}

func _error(){
	user, err := getUser("")

	if errors.Is(err, UserNotFoundError) {
		fmt.Println("User not found error occurred")
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}

	err = validateUserInput("")
	if err != nil {
		fmt.Println("Validation Error:", err)
	} else {
		fmt.Println("Input is valid")
	}
}