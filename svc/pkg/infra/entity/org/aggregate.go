package entity

import "ynufes-mypage-backend/svc/pkg/domain/model/org"

type Org struct {
	ID      org.ID  `firestore:"-"`
	EventID int64   `firestore:"event_id"`
	Name    string  `firestore:"name"`
	Members []int64 `firestore:"member_ids"`
}
