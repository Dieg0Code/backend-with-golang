package controllers

import (
	"net/http"
	"strconv"

	"github.com/dieg0code/go-microservices/services"
	"github.com/dieg0code/go-microservices/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		utils.RespondError(c, apiErr)
		// c.JSON(apiErr.Status, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		// c.JSON(apiErr.Status, apiErr)
		return
	}

	utils.Respong(c, http.StatusOK, user)
	// c.JSON(http.StatusOK, user)

}
