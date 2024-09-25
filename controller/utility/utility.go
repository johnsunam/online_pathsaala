package utility

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field,omitempty"`
	Msg   string `json:"message"`
}

func CovertValidationErrorMsg(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			fmt.Println(fe.Param())
			out[i] = ApiError{fe.Field(), msgForTag(fe.Tag(), fe.Param())}
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": out})
		return
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": []ApiError{{Msg: "invalid payload format"}}})
}

func ConvertErrorMessage(c *gin.Context, err error, statusCode int) {
	errMessage := ApiError{Msg: err.Error()}
	c.JSON(statusCode, gin.H{"errors": []ApiError{errMessage}})
}

func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("Field should contain atleast %s characters", param)
	case "eqfield":
		return "password should match the confirm password"
	case "oneof":
		return fmt.Sprintf("Field should have one of the value %s", param)
	default:
		return "Invalid field value provided"
	}
}
