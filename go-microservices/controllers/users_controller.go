package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dieg0code/go-microservices/services"
	"github.com/dieg0code/go-microservices/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {

	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}
