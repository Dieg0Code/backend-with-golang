package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequest(t *testing.T) {

	request := CreateRepoRequest{
		Name:        "golang intro",
		Description: "a golang intro repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	// Convert struct to JSON
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest

	// Convert JSON to struct
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	// Compare the original struct with the converted struct to make sure they are the same
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProjects, request.HasProjects)
	assert.EqualValues(t, target.HasWiki, request.HasWiki)
}
