package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK , "signUp")
}