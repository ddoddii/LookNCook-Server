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
