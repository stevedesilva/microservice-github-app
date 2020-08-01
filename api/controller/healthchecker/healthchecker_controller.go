package healthchecker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	up = "UP"
)

// CheckUp is heath checker
func CheckUp(c *gin.Context) {

	c.String(http.StatusOK, up)
}
