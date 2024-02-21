package handler

import (
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"github.com/gdsc-ys/lookncook-server/src/model"
	"github.com/gdsc-ys/lookncook-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RecipeHandler interface {
	GetRecipe(e echo.Context) error
}

type RecipeHandlerImpl struct {
	recipeService service.RecipeService
}

func (handler *RecipeHandlerImpl) GetRecipe(c echo.Context) error {
	imgData, err := utils.ExtractImageData(c, "image")
	if err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, "Error with image"+err.Error())
	}
	fridgeContent, recipeContent := handler.recipeService.GenerateRecipeFromFridgeContent(imgData)

	return c.JSON(http.StatusOK, map[string]string{
		"fridge": fridgeContent,
		"recipe": recipeContent,
	})
}

func RecipeHandlerInit(recipeService service.RecipeService) *RecipeHandlerImpl {
	return &RecipeHandlerImpl{
		recipeService: recipeService,
	}
}
