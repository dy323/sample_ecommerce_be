package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ecommerce/controller"
	"ecommerce/auth"
	"ecommerce/queue"
)

func main() {
	//ready queue jobs
	queue.StartQueue()

	e := echo.New()

	g := e.Group("/api/admin")

	g.Use(auth.AuthVerification)

	e.Use(middleware.Recover()) 

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
		
	e.GET("/api/index", controller.GetIndex)

	e.GET("/api/category", controller.GetCategory)

	e.GET("/api/product/:id", controller.GetProduct)

	e.GET("/api/verify/:id", controller.Verify)

	e.POST("/api/search", controller.GetSearch)

	e.POST("/api/login", controller.Login)

	e.POST("/api/signup", controller.Register)

	g.POST("/profile", controller.GetProfile)

	e.Logger.Fatal(e.Start(":8000"))
}