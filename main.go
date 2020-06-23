package main

import (
	//This should be the docs file where swag generates its files
	_ "echoSwaggerTest/docs"
	"echoSwaggerTest/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// wrapper for swagger
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger POC
// @version 1
// @description Just a simple poc for swagger
// @termsOfService http://swagger.io/terms/

// @contact.name Cassio Roos
// @contact.url https://github.com/CassioRoos/
// @contact.email cassio.roos@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7777
func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())

	//Routes
	e.GET("/hello", handlers.Hello)
	e.GET("/goodnight/:to", handlers.GetSayGoodNigth)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start server
	e.Logger.Fatal(e.Start(":7777"))
}
