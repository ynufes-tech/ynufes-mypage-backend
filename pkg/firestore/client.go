package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"os"
	"ynufes-mypage-backend/pkg/setting"
)

var (
	Client *firestore.Client
)

func init() {
	ctx := context.Background()
	config := setting.Get()
	data, err := os.ReadFile(config.Infrastructure.Firestore.JsonCredentialFile)
	options := option.WithCredentialsJSON(data)
	Client, err = firestore.NewClient(ctx, config.Infrastructure.Firestore.ProjectID, options)
	if err != nil {
		panic(err)
	}
}
