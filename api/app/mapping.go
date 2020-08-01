package app

import (
	"github.com/stevedesilva/golang-microservices/src/api/controller/healthchecker"
	repositories "github.com/stevedesilva/golang-microservices/src/api/controller/respositories"
)

func mapUrls() {
	router.GET("/health", healthchecker.CheckUp)
	router.POST("/repositories", repositories.CreateRepo)
}
