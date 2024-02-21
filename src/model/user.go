package model

type User struct {
	Id   string
	Name string
	Age  int
}

type UserDevice struct {
	FcmToken string
	DeviceId string
}

type UserDeviceRequest struct {
	DeviceId string `json:"DeviceId"`
}
