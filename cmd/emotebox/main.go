package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yeno-team/emotebox/controller"
	_ "github.com/yeno-team/emotebox/docs"
	"github.com/yeno-team/emotebox/pkg/example"
)

// @title Emotebox API
// @version 1.0
// @description Emotebox's Server

// @license.name MIT
// @license.url https://github.com/yeno-team/emotebox/blob/main/LICENSE

func main() {
	r := gin.Default()

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
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
