package handler

import (
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"github.com/gdsc-ys/lookncook-server/src/model"
	"github.com/gdsc-ys/lookncook-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type KitchenHandler interface {
	GetKitchenEnvironment(e echo.Context) error
}

type KitchenHandlerImpl struct {
	kitchenService service.KitchenService
}

// GetKitchenEnvironment godoc
// @Summary Get kitchen environment from an image
// @Description Upload an image to analyze and get a description of the kitchen environment.
// @Tags kitchen
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "image"
// @Success 200 {object} model.KitchenEnvironmentResponse
// @Failure 400 {object} model.CustomHTTPError "Error with image"
// @Failure 500 {object} model.CustomHTTPError "Internal server error"
// @Router /GetKitchenEnvironment [post]
func (handler *KitchenHandlerImpl) GetKitchenEnvironment(c echo.Context) error {
	imgData, err := utils.ExtractImageData(c, "image")
	if err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, "Error with image")
	}

	content, err := handler.kitchenService.GetCookingEnvironment(imgData)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"content": content,
	})
}

func KitchenHandlerInit(kitchenService service.KitchenService) *KitchenHandlerImpl {
	return &KitchenHandlerImpl{
		kitchenService: kitchenService,
	}
}
