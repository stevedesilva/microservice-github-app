package github_provider_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stevedesilva/golang-microservices/src/api/client/restclient"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/golang-microservices/src/api/domain/github"
	. "github.com/stevedesilva/golang-microservices/src/api/providers/github_provider"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := "123"
	actual := GetAuthorizationHeader(header)
	assert.EqualValues(t, "token 123", actual)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.FlushMockup()
	restclient.AddMockup(restclient.Mock{
		Response:   nil,
		Err:        errors.New("invalid rest client response"),
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid rest client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockup()

	invalidCloser, _ := os.Open("aaa")
	restclient.AddMockup(restclient.Mock{
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockup()

	reader := strings.NewReader(`{"message": 1}`)
	readCloser := ioutil.NopCloser(reader)

	restclient.AddMockup(restclient.Mock{
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       readCloser,
		},
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockup()

	reader := strings.NewReader(`{
		"message": "Requires authentication",
		"documentation_url": "https://developer.github.com/v3"
	}`)
	readCloser := ioutil.NopCloser(reader)

	restclient.AddMockup(restclient.Mock{
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       readCloser,
		},
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockup()

	reader := strings.NewReader(`{"id": "1234"}`)
	readCloser := ioutil.NopCloser(reader)

	restclient.AddMockup(restclient.Mock{
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       readCloser,
		},
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal create repo successful response", err.Message)
}

func TestCreateRepoHappy(t *testing.T) {
	restclient.FlushMockup()

	reader := strings.NewReader(`{
		"id": 1234,
		"name": "golang-tutorial",
		"full_name": "stevedesilva/golang-tutorial"}`)
	readCloser := ioutil.NopCloser(reader)

	restclient.AddMockup(restclient.Mock{
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       readCloser,
		},
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	res, err := CreateRepo("accesstoken", github.CreateRepoRequest{})
	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.EqualValues(t, 1234, res.ID)
	assert.EqualValues(t, "golang-tutorial", res.Name)
	assert.EqualValues(t, "stevedesilva/golang-tutorial", res.FullName)
}

func TestCreateRepoConstants(t *testing.T) {
	assert.EqualValues(t,"Authorization",HeaderAuthorization)
	assert.EqualValues(t,"token %s",HeaderAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos",UrlCreateRepo)
}