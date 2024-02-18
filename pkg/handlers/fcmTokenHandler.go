package handlers

import (
	"github.com/labstack/echo/v4"
	"lookncook/pkg/services"
	"net/http"
)

type FcmTokenRequest struct {
	DeviceId string `json:"deviceId"`
	FcmToken string `json:"fcmToken"`
}

func FcmTokenHandler(c echo.Context) error {
	var req FcmTokenRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	services.FcmTokenService(req.DeviceId, req.FcmToken)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "FCM token added successfully",
	})
}
