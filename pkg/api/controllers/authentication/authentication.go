package authentication

import (
	"net/http"
	errorHandler "urlShortner/pkg/errors"
	"urlShortner/pkg/models/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate *validator.Validate

func SignUp(c *gin.Context) {
	validate = validator.New()

	var newUser user.User
	err := c.ShouldBind(&newUser)

	// if err != nil {
	// 	c.IndentedJSON(http.StatusOK , err)
	// 	return
	// }
		
	// ValidationErr := validate.Struct(newUser)

	if err != nil {
		log.Errorln(errorHandler.Handler(err , c))
		// c.IndentedJSON(http.StatusOK , gin.H{"error" :errorHandler.Handler(err)})
		return
	}

	newUser.InitUser()
	// log.Info(newUser)

	c.IndentedJSON(http.StatusOK , "signUp")
	return
}