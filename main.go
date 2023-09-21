package main

import (
	"scheduler/config"
	"scheduler/controllers"
	"scheduler/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitMongoDB()

	e := echo.New()

	e.POST("/transactions", controllers.CreateTransaction)
	e.GET("/transactions", controllers.GetAllTransaction)
	e.GET("/transactions/:id", controllers.GetTransaction)
	e.PUT("/transactions/:id", controllers.UpdateTransaction)
	e.DELETE("/transactions/:id", controllers.DeleteTransaction)

	services.StartScheduler()

	e.Logger.Fatal(e.Start(":8080"))
}