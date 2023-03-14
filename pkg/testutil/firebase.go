package testutil

import (
	"context"
	fb "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"log"
	"ynufes-mypage-backend/pkg/firebase"
)

type FirebaseTest struct {
	db *db.Client
}

func NewFirebaseTest() *FirebaseTest {
	ctx := context.Background()
	conf := &fb.Config{
		DatabaseURL: "localhost:9000/?ns=ynufes-mypage-default-rtdb",
	}
	app, err := fb.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}
	c, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	return &FirebaseTest{
		db: c,
	}
}

func (f FirebaseTest) GetClient() *firebase.Firebase {
	fdb := firebase.NewWithClient(f.db)
	return &fdb
}

func (f FirebaseTest) Reset() {
	err := f.db.NewRef("/").Delete(context.Background())
	if err != nil {
		log.Println("failed to reset firebase: ", err)
		return
	}
}
