package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeno-team/emotebox"
)

type UserController struct {
	UserService emotebox.UserService
}

// User Get Controller godoc
// @Summary Gets the current authenticated user
// @Description Gets the current authenticated user
// @Tags Authentication
// @Success 200 {object} emotebox.User
// @Router /api/v1/user [get]
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	user, err := uc.UserService.GetCurrentUser()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}
