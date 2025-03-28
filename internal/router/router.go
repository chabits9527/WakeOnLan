package router

import (
	"WakeOnLan/internal/controller"
	"WakeOnLan/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(wakeController *controller.WakeController) *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.Logger) // 添加日志中间件
	r.Use(middleware.RecoveryMiddleware())
	r.Group("/api")
	{
		r.GET("/wake/onLAN", wakeController.WakeOnLAN)
	}
	return r
}
