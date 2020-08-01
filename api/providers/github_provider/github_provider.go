package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stevedesilva/golang-microservices/src/api/client/restclient"
	"github.com/stevedesilva/golang-microservices/src/api/domain/github"
)

const (
	HeaderAuthorization       = "Authorization"
	HeaderAuthorizationFormat = "token %s"
	UrlCreateRepo             = "https://api.github.com/user/repos"
)

// GetAuthorizationHeader function
func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(HeaderAuthorizationFormat, accessToken)
}

// CreateRepo function
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	header := http.Header{}
	header.Set(HeaderAuthorization, GetAuthorizationHeader(accessToken))
	fmt.Println("header: ", header)

	response, err := restclient.Post(UrlCreateRepo, request, header)
	fmt.Println("response: ", response)
	fmt.Println("error: ", err)
	if err != nil {
		log.Println(fmt.Printf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal create repo successful response"}
	}

	return &result, nil
}
