package services

import (
	"github.com/dieg0code/go-microservices/domain"
	"github.com/dieg0code/go-microservices/utils"
)

type itemService struct{}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "Not implemented",
		Status:  500,
		Code:    "internal_server_error",
	}
}
