package services

import (
	"testing"

	"github.com/dieg0code/go-microservices/domain"
	"github.com/dieg0code/go-microservices/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock
)

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return &domain.User{Id: 123, FirstName: "Diego", LastName: "Obando", Email: "example@email.com"}, nil
}

func TestGetUserNotUserFoundInDatabase(t *testing.T) {

	user, err := UserService.GetUser(0)
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, 404, err.Status, "We were expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting 'not_found' error code when user is not found")
	assert.EqualValues(t, "User 0 was not found", err.Message, "We were expecting 'User 0 was not found' error message when user is not found")

}
