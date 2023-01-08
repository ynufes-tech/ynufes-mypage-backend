package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type Org struct {
	ID      org.ID  `firestore:"-"`
	EventID int64   `firestore:"event_id"`
	Name    string  `firestore:"name"`
	Members []int64 `firestore:"member_ids"`
}

func (o Org) ToModel() (*org.Org, error) {
	var members []user.ID
	for i := range o.Members {
		members = append(members, identity.NewID(o.Members[i]))
	}
	return &org.Org{
		ID:      o.ID,
		EventID: identity.NewID(o.EventID),
		Name:    o.Name,
		Members: members,
	}, nil
}
