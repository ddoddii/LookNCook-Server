package handler

import (
	"fmt"
	"github.com/gdsc-ys/lookncook-server/src/model"
	"github.com/gdsc-ys/lookncook-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler interface {
	GetUserDevice(c echo.Context) error
	AddUserDevice(c echo.Context) error
}

type UserHandlerImpl struct {
	userService service.UserService
}

func (handler *UserHandlerImpl) GetUserDevice(c echo.Context) error {
	var userDeviceRequest model.UserDeviceRequest

	if err := c.Bind(&userDeviceRequest); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, "Invalid Request Body")
	}

	if userDeviceRequest.DeviceId == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "deviceID is required")
	}

	fmt.Println(userDeviceRequest.DeviceId)

	userDevice, err := handler.userService.GetUserDevice(userDeviceRequest.DeviceId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userDevice)
}

func (handler *UserHandlerImpl) AddUserDevice(c echo.Context) error {
	var userDevice model.UserDevice

	if err := c.Bind(&userDevice); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, "Invalid form data")
	}

	_, err := handler.userService.AddUserDevice(userDevice.DeviceId, userDevice.FcmToken)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": fmt.Sprintf("%s added successfully", userDevice.DeviceId)})
}

func UserHandlerInit(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}
