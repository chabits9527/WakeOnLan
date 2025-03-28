package main

import (
	"WakeOnLan/internal/config"
	"WakeOnLan/internal/controller"
	"WakeOnLan/internal/router"
	"WakeOnLan/internal/service"
	"fmt"
)

func main() {
	appConfig := config.Load()
	wakeService := service.CreateWakeService(appConfig)
	wakeController := controller.CreateWakeController(wakeService)

	r := router.NewRouter(wakeController)
	r.Run(fmt.Sprintf(":%d", appConfig.Port))
}
