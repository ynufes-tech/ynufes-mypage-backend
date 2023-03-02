package infra

import (
	"cloud.google.com/go/firestore"
	"context"
	"testing"
	"ynufes-mypage-backend/pkg/testutil"
)

var client *firestore.Client

//func TestFirestore(t *testing.T) {
//	w := writer.NewUser(client)
//	test1 := testutil.Users()
//	err := w.Create(context.Background(), test1[0])
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.IsEqual(err, nil)
//
//	r := reader.NewUser(client, )
//	u, err := r.GetByID(context.Background(), id.UserID(1234))
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(u)
//}

func TestMain(m *testing.M) {
	c, killer := testutil.NewFirestoreTestClient(context.Background())
	client = c
	if killer != nil {
		defer killer()
	}
	m.Run()
}
