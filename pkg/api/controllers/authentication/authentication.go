package authentication

import (
	"fmt"
	"net/http"
	"urlShortner/pkg/models/user"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser user.User

	err := c.BindJSON(&newUser)

	if err != nil {
		c.IndentedJSON(http.StatusOK , err)
		return
	}

	newUser.InitUser()
	fmt.Println(newUser)

	c.IndentedJSON(http.StatusOK , "signUp")
}