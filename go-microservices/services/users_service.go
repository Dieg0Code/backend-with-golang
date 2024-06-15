package services

import (
	"github.com/dieg0code/go-microservices/domain"
	"github.com/dieg0code/go-microservices/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
