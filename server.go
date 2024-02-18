package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"lookncook/internal/config"
	"lookncook/pkg/handlers"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.POST("/fridge", handlers.FridgeContentHandler)
	e.POST("/fridge-recipe", handlers.RecipeHandler)
	e.POST("/fcmToken", handlers.FcmTokenHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
