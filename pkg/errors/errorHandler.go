package errorHandler

import (
	"errors"

	customErrors "urlShortner/pkg/models/errors"

	"github.com/go-playground/validator/v10"
)


func messageForError(err validator.FieldError) string {
	switch err.Tag() {
	case "required" :
		return "This Field is required"
	}

	return err.Error()
}

func Handler(err error) []customErrors.ApiError {

	var ve validator.ValidationErrors
	var out []customErrors.ApiError

	if errors.As(err , &ve) {
		out := make([]customErrors.ApiError , len(ve))
		for i,fe := range ve {
			out[i] = customErrors.ApiError{Param : fe.Field() ,Message : messageForError(fe)}
		}
	}

	// fmt.Print(out)

	return out	
}