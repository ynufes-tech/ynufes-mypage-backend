package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"log"
	"os"
	"ynufes-mypage-backend/pkg/setting"
)

var (
	client *firestore.Client
)

func New() *firestore.Client {
	if client == nil {
		ctx := context.Background()
		config := setting.Get()
		if config.Infrastructure.Firestore.JsonCredentialFile == "TESTING" {
			client, _ = firestore.NewClient(ctx, "ynufes-mypage")
			return client
		}
		data, err := os.ReadFile(config.Infrastructure.Firestore.JsonCredentialFile)
		options := option.WithCredentialsJSON(data)
		client, err = firestore.NewClient(ctx, config.Infrastructure.Firestore.ProjectID, options)
		if err != nil {
			log.Fatalf("firebase.NewClient err: %v", err)
		}
	}
	return client
}
