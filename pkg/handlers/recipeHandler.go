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
	promptText, err := utils.ReadFileAsString("internal/prompt/recipePrompt.txt")

	content := services.GeminiImageToTextFunction(imgData, promptText)

	return c.JSON(http.StatusOK, content)
}
