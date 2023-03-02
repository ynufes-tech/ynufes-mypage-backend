package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
	"log"
	"ynufes-mypage-backend/pkg/setting"
)

type Firebase struct {
	client *db.Client
}

func New() Firebase {
	config := setting.Get()
	if config.Infrastructure.Firebase.DatabaseURL == "" {
		panic("firebase database url is not set")
	}
	if config.Infrastructure.Firebase.JsonCredentialFile == "" {
		panic("firebase credential path is not set")
	}
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: config.Infrastructure.Firebase.DatabaseURL,
	}
	opt := option.WithCredentialsFile(
		config.Infrastructure.Firebase.JsonCredentialFile,
	)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	c, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	return Firebase{
		client: c,
	}
}

func (f Firebase) Client(path string) *db.Ref {
	return f.client.NewRef(path)
}
