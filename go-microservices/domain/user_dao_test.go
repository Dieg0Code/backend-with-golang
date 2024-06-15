package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.Status, "We were expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting 'not_found' error code when user is not found")
	assert.EqualValues(t, "User 0 was not found", err.Message, "We were expecting 'User 0 was not found' error message when user is not found")
}

func TestGerUserNoError(t *testing.T) {

	user, err := UserDao.GetUser(123)

	assert.Nil(t, err, "We were not expecting an error when user id is 123")
	assert.NotNil(t, user, "We were expecting a user with id 123")
	assert.EqualValues(t, 123, user.Id, "We were expecting a user with id 123")
	assert.EqualValues(t, "Diego", user.FirstName, "We were expecting a user with first name Diego")
	assert.EqualValues(t, "Obando", user.LastName, "We were expecting a user with last name Obando")
	assert.EqualValues(t, "example@email.com", user.Email, "We were expecting a user with email")
}
