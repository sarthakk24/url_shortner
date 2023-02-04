package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK , "working")
}