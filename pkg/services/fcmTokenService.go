package services

import "lookncook/internal/database"

type FcmTokenRequest struct {
	DeviceId string `json:"deviceId"`
	FcmToken string `json:"fcmToken"`
}

// Add deviceId & fcmToken to Firestore

func FcmTokenService(deviceId string, fcmToken string) {
	client := database.InitFirebase()
	database.AddDeviceIdAndFcmToken(deviceId, fcmToken, client)
}
