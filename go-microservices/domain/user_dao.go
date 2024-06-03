package domain

import (
	"fmt"
	"net/http"

	"github.com/dieg0code/go-microservices/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "example@email.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
