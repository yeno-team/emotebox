/**
 * Copyright (C) 2022 Yeno
 *
 * This file is part of Emotebox.
 *
 * Emotebox is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Emotebox is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Emotebox.  If not, see <http://www.gnu.org/licenses/>.
 */

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
