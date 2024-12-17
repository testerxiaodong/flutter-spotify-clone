package controllers

import (
	"server/common/response"
	"server/common/xerror"
	"server/services"
	"server/types"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (ctrl *UserController) Signup(c *gin.Context) {
	var signupReq types.SignupReq
	if err := c.ShouldBindJSON(&signupReq); err != nil {
		response.Result(c, nil, xerror.NewErrCodeMsg(xerror.REQUEST_PARAM_ERROR, err.Error()))
		return
	}
	res, err := ctrl.UserService.Signup(&signupReq)
	if err != nil {
		response.Result(c, nil, err)
		return
	}
	response.Result(c, res, nil)
}
