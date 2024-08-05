package main

import (
	"github.com/labstack/echo/v4"
	"go-nabati/controllers/authcontroller"
	"go-nabati/controllers/redemptioncontroller"
	"go-nabati/controllers/usercontroller"
	"go-nabati/middlewares"
	"go-nabati/models"
)

func main() {
	e := echo.New()
	models.ConnectDB()

	//auth
	e.POST("/login", authcontroller.Login)
	e.GET("/logout", authcontroller.Logout)
	//user
	e.GET("/users", usercontroller.Index, middlewares.JWTMiddleware)
	e.GET("/users/:id", usercontroller.Show, middlewares.JWTMiddleware)
	e.POST("/users", usercontroller.Create, middlewares.JWTMiddleware)
	e.PUT("/users/:id", usercontroller.Update, middlewares.JWTMiddleware)
	e.DELETE("/users/:id", usercontroller.Delete, middlewares.JWTMiddleware)
	//unique code
	e.GET("/code", redemptioncontroller.Index)
	e.POST("/code/check", redemptioncontroller.Check)
	e.POST("/code/redeem", redemptioncontroller.Submit)
	e.Logger.Fatal(e.Start(":1323"))
}
