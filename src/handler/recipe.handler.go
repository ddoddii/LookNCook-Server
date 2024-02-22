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

// GetRecipe godoc
// @Summary Get a recipe based on a provided fridge image
// @Description Upload an image of your fridge's contents to receive a recipe recommendation.
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "image"
// @Success 200 {object} model.RecipeResponse "Successfully generated recipe"
// @Failure 400 {object} model.CustomHTTPError "Error with the provided image"
// @Router /GetFridgeRecipe [post]
func (handler *RecipeHandlerImpl) GetRecipe(c echo.Context) error {
	imgData, err := utils.ExtractImageData(c, "image")
	if err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, "Error with image "+err.Error())
	}
	ingredients, recipes := handler.recipeService.GenerateRecipeFromFridgeContent(imgData)

	response := model.RecipeResponse{
		Ingredients: ingredients,
		Recipes:     recipes,
	}

	return c.JSON(http.StatusOK, response)
}

func RecipeHandlerInit(recipeService service.RecipeService) *RecipeHandlerImpl {
	return &RecipeHandlerImpl{
		recipeService: recipeService,
	}
}
