package controller

import (
	"WakeOnLan/internal/service"
	"WakeOnLan/internal/utils/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WakeController struct {
	wakeService *service.WakeService
}

func CreateWakeController(s *service.WakeService) *WakeController {
	return &WakeController{
		wakeService: s,
	}
}

func (c *WakeController) WakeOnLAN(ctx *gin.Context) {
	c.wakeService.WakeOnLAN()
	ctx.JSON(http.StatusOK, result.DefaultSuccess())
}
