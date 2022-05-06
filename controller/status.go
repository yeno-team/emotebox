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
