package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/gdsc-ys/lookncook-server/src/constant"
	"github.com/gdsc-ys/lookncook-server/src/model"
)

type UserService interface {
	AddUserDevice(deviceId string, fcmToken string) (model.UserDevice, error)
	GetUserDevice(deviceId string) (model.UserDeviceResponse, error)
}

type UserServiceImpl struct {
	firestoreClient *firestore.Client
}

func (service *UserServiceImpl) AddUserDevice(deviceId string, fcmToken string) (model.UserDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	userDevice := model.UserDevice{
		DeviceId: deviceId,
		FcmToken: fcmToken,
	}

	fmt.Println(userDevice)

	_, _, err := service.firestoreClient.
		Collection(constant.UserDeviceCollection).
		Add(ctx, userDevice)

	if err != nil {
		return model.UserDevice{}, err
	}
	return userDevice, nil

}

func (service *UserServiceImpl) GetUserDevice(deviceId string) (model.UserDeviceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	query := service.firestoreClient.
		Collection(constant.UserDeviceCollection).
		Where("DeviceId", "==", deviceId).
		Limit(1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		return model.UserDeviceResponse{}, err
	}

	var userDevice = model.UserDeviceResponse{}
	if err := doc.DataTo(&userDevice); err != nil {
		return model.UserDeviceResponse{}, err
	}

	return userDevice, nil
}

func UserServiceInit(firestoreClient *firestore.Client) *UserServiceImpl {
	return &UserServiceImpl{
		firestoreClient: firestoreClient,
	}
}
