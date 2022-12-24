package infra

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
	"ynufes-mypage-backend/svc/pkg/infra/writer"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

func TestFirestore(t *testing.T) {
	w := writer.NewUser(client)
	test1 := testutil.Users()
	err := w.Create(context.Background(), test1[0])
	if err != nil {
		t.Fatal(err)
	}
	assert.IsEqual(err, nil)

	r := reader.NewUser(client)
	u, err := r.GetByID(context.Background(), 1234)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestMain(m *testing.M) {
	c, killer := testutil.NewFirestoreTestClient(context.Background())
	client = c
	if killer != nil {
		defer killer()
	}
	m.Run()
}
