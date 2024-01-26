package handlers

import (
	"github.com/labstack/echo/v4"
	"lookncook/internal/utils"
	"lookncook/pkg/services"
	"net/http"
)

func FridgeContentHandler(c echo.Context) error {
	imgData, err := utils.ExtractImageData(c, "image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error with image: " + err.Error()})
	}

	content := services.FridgeContentGenFunction(imgData)

	return c.JSON(http.StatusOK, content)
}
