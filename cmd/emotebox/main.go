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

package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yeno-team/emotebox/controller"
	_ "github.com/yeno-team/emotebox/docs"
	"github.com/yeno-team/emotebox/middleware"
	"github.com/yeno-team/emotebox/pkg/discordgo"
	"github.com/yeno-team/emotebox/pkg/example"
)

// @title Emotebox API
// @version 1.0
// @description Emotebox's Server

// @license.name MIT
// @license.url https://github.com/yeno-team/emotebox/blob/main/LICENSE

func main() {
	APP_ENV := strings.ToLower(os.Getenv("APP_ENV"))
	if APP_ENV == "" {
		APP_ENV = "development"
	}
	err := godotenv.Load("." + APP_ENV + ".env")
	if err != nil {
		panic(err)
	}

	logger := &example.Logger{}

	discordSession, err := discordgo.New(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.ErrorHandler(logger))

	v1 := r.Group("/api/v1")
	{
		status := v1.Group("/status")
		{
			statusService := &example.StatusService{}
			c := &controller.StatusController{
				StatusService: statusService,
			}
			status.GET("", c.Default)
		}

		user := v1.Group("/user")
		{
			userService := &discordgo.UserService{Session: discordSession}
			c := &controller.UserController{
				UserService: userService,
			}
			user.GET("", c.GetCurrentUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
