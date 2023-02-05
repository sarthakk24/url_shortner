package errorHandler

import (
	"errors"
	"fmt"
	"net/http"

	customErrors "urlShortner/pkg/models/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func messageForError(err string) string {
	switch err {
	case "required" :
		return "This Field is required"
		
	case "email":
		return "Email is not valid"
	}

	
	return "test"
}

func Handler(err error , c *gin.Context) []customErrors.ApiError {
	var ve validator.ValidationErrors
	var out []customErrors.ApiError

	if errors.As(err , &ve) {
		out := make([]customErrors.ApiError , len(ve))
		for i,fe := range ve {
			out[i] = customErrors.ApiError{Param : fe.Field() ,Message : messageForError(fe.Tag())}

		}
		
		c.JSON(http.StatusAccepted , gin.H{"errors" : out})
	}

	fmt.Println(out)

	return out	
}