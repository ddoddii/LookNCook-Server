package handlers

import (
	"github.com/labstack/echo/v4"
	"lookncook/internal/utils"
	"lookncook/pkg/services"
	"net/http"
)

func RecipeHandler(c echo.Context) error {
	imgData, err := utils.ExtractImageData(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error with image: " + err.Error()})
	}
	fridgeContent, recipeContent := services.GenerateRecipeFromFridgeContent(imgData)

	return c.JSON(http.StatusOK, map[string]string{
		"fridge": fridgeContent,
		"recipe": recipeContent,
	})
}
