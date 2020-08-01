package config

import (
	"fmt"
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetAccessToken() string {
	fmt.Println("githubAccessToken: ", githubAccessToken)
	return githubAccessToken
}
