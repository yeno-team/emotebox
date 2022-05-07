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

type StatusController struct {
	StatusService emotebox.StatusService
}

// Default godoc
// @Summary Gives the default status of the app
// @Description get status of app
// @Tags Status
// @Success 200 {object} emotebox.Status
// @Router /api/v1/status [get]
func (sc *StatusController) Default(c *gin.Context) {
	defaultStatus, err := sc.StatusService.Status()
	if err != nil {
		// panic since this should never happen
		panic(err)
	}
	c.JSON(defaultStatus.Code, defaultStatus)
}
