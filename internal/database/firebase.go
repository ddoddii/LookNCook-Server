package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

func InitFirebase() *firestore.Client {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Success in connecting")
	return client
}

func AddDeviceIdAndFcmToken(deviceId string, fcmToken string, client *firestore.Client) {
	ctx := context.Background()
	_, _, err := client.Collection("fcmTokens").Add(ctx, map[string]interface{}{
		"deviceId": deviceId,
		"fcmToken": fcmToken,
	})
	fmt.Println("Success in adding fcmToken and deviceId")
	if err != nil {
		log.Fatalf("Failed adding new fcm token: %v", err)
	}
}

func ViewFcmTokens(client *firestore.Client) {
	ctx := context.Background()
	iter := client.Collection("fcmTokens").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
