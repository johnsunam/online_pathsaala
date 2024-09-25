package user

import (
	"context"
	"errors"
	"net/http"

	"online-pathsaala/controller/utility"
	"online-pathsaala/model"
	"online-pathsaala/pkg/db"

	"github.com/gin-gonic/gin"
)

type DbInterface interface {
	Register(context.Context, model.RegisterPayload) (string, error)
	GetUser(context.Context, model.LoginPayload) (model.User, error)
}

type UserAcc struct {
	DdManager DbInterface
}

func (usr *UserAcc) Register(c *gin.Context) {
	var reqPayload model.RegisterPayload
	ctx := c.Request.Context()

	if err := c.ShouldBindBodyWithJSON(&reqPayload); err != nil {
		utility.CovertValidationErrorMsg(c, err)
		return
	}

	_, err := usr.DdManager.Register(ctx, reqPayload)
	if err != nil {
		utility.ConvertErrorMessage(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, "User signed up successfully")
}

func (usr *UserAcc) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var reqPayload model.LoginPayload
	if err := c.ShouldBindBodyWithJSON(&reqPayload); err != nil {
		utility.CovertValidationErrorMsg(c, err)
		return
	}
	user, err := usr.DdManager.GetUser(ctx, reqPayload)
	if err != nil {
		utility.ConvertErrorMessage(c, err, http.StatusInternalServerError)
		return
	}
	if user.ID == "" {
		utility.ConvertErrorMessage(c, errors.New("user doesn't exist in system"), http.StatusNotFound)
		return
	}
	check, msg := db.VerifyPassword(user.Password, reqPayload.Password)
	if !check {
		utility.ConvertErrorMessage(c, errors.New(msg), http.StatusUnauthorized)
		return
	}
	token, err := db.GenerateTokens(user.Email, user.ID)
	if err != nil {
		utility.ConvertErrorMessage(c, errors.New("internal server error while creating jwt token"), http.StatusUnprocessableEntity)
		return
	}
	response := model.LoginResponse{
		Id:    user.ID,
		Email: user.Email,
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}
