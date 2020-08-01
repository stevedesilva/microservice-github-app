package services

import (
	"strings"

	"github.com/stevedesilva/golang-microservices/src/api/config"

	"github.com/stevedesilva/golang-microservices/src/api/providers/github_provider"

	"github.com/stevedesilva/golang-microservices/src/api/domain/github"

	r "github.com/stevedesilva/golang-microservices/src/api/domain/repositories"
	"github.com/stevedesilva/golang-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(input r.CreateRepoRequest) (*r.CreateRepoResponse, errors.ApiError)
}

var (
	// RepositoryService exposed to client - only way to get to the service
	RepositoryService reposServiceInterface
)

func init() {
	// An interface can store either a struct directly or a pointer to a struct.
	RepositoryService = &reposService{}
}


func (s *reposService) CreateRepo(input r.CreateRepoRequest) (*r.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	response, err := github_provider.CreateRepo(config.GetAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := &r.CreateRepoResponse{
		Id:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return result, nil
}
