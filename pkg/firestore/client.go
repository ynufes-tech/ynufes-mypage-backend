package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
)

var (
	Client *firestore.Client
)

func init() {
	ctx := context.Background()
	var err error
	Client, err = firestore.NewClient(ctx, "ynufes-mypage")
	if err != nil {
		panic(err)
	}
}
