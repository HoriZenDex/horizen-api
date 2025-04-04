package main

import (
	"log"

	"horizen-api/config"
	"horizen-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Create a new Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup routes, including those that require JWT authentication
	routes.RegisterRoutes(e)

	// Start the server on port 8080
	log.Fatal(e.Start(":8080"))
}
