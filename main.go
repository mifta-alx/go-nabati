package main

import (
	"github.com/labstack/echo/v4"
	"go-nabati/controllers/usercontroller"
	"go-nabati/models"
)

func main() {
	e := echo.New()
	models.ConnectDB()
	e.GET("/users", usercontroller.Index)
	e.GET("/user/:id", usercontroller.Show)
	e.POST("/user", usercontroller.Create)
	e.PUT("/user/:id", usercontroller.Update)
	e.DELETE("/user", usercontroller.Delete)
	e.Logger.Fatal(e.Start(":1323"))
}
