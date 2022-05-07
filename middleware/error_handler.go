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

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yeno-team/emotebox"
)

func ErrorHandler(logger emotebox.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) <= 0 {
			return
		}

		for _, ginErr := range c.Errors {
			logger.Warnf("ERROR: ", ginErr.Err.Error())
		}

		c.JSON(-1, gin.H{
			"errors": c.Errors,
		})
	}
}
