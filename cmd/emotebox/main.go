package main

import (
	"os"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	logger := &example.Logger{}
	r.Use(middleware.ErrorHandler(logger))
	discordSession, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic(err)
	}

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
