package services

import (
	"github.com/dieg0code/go-microservices/domain"
	"github.com/dieg0code/go-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
