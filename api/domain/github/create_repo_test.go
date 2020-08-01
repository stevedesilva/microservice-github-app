package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	/*
	   Test Marshalling take input interface and attempt to create a json string ( obj -> json )
	*/
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introductory respository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)      // not error in marshalling
	assert.NotNil(t, bytes) // we have json string

	/*
	   Test Unmarshalling
	*/
	var target CreateRepoRequest         // struct to populate
	err = json.Unmarshal(bytes, &target) // json string to target struct

	// compare to original request
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.True(t, target.Private)
	assert.True(t, target.HasProjects)
	assert.True(t, target.HasIssues)
	assert.True(t, target.HasWiki)
}

// TODO test response
/*
type CreateRepoResponse struct {
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	FullName   string          `json:"full_name"`
	Owner      RepoOwner       `json:"owner"`
	Permssions RepoPermissions `json:"permssions"`
}
*/
func TestCreateRepoResponseAsJson(t *testing.T) {
	// marshalling
	response := CreateRepoResponse{
		ID:          1,
		Name:        "Steve",
		FullName:    "Steve deSilva",
		Owner:       RepoOwner{ID: 1, Login: "login", Url: "url", HtmlUrl: "htmlUrl"},
		Permissions: RepoPermissions{IsAdmin: true, IsPush: true, IsPull: true},
	}

	bytes, err := json.Marshal(response)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// Unmarshalling
	var target CreateRepoResponse
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)

	assert.EqualValues(t, response.ID, target.ID)
	assert.EqualValues(t, response.Name, target.Name)
	assert.EqualValues(t, response.FullName, target.FullName)
	assert.EqualValues(t, response.Owner, target.Owner)
	assert.EqualValues(t, response.Permissions, target.Perimssions)
}
