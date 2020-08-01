package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevedesilva/golang-microservices/src/api/domain/repositories"
	"github.com/stevedesilva/golang-microservices/src/api/services"
	"github.com/stevedesilva/golang-microservices/src/api/utils/errors"
)

// CreateRepo func
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	resp, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
