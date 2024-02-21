package router

import (
	"github.com/gdsc-ys/lookncook-server/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Router(init *config.Initialization) *echo.Echo {
	e := echo.New()

	e.Debug = true

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/GetFridgeRecipe", init.RecipeHandler.GetRecipe)
	e.POST("/GetKitchenEnvironment", init.KitchenHandler.GetKitchenEnvironment)

	e.POST("/StoreUserDevice", init.UserHandler.AddUserDevice)
	e.GET("/GetFcmToken", init.UserHandler.GetUserDevice)

	return e
}
