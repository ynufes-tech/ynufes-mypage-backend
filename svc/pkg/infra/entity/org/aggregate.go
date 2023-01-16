package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type Org struct {
	ID        string  `firestore:"-"`
	EventID   int64   `firestore:"event_id"`
	EventName string  `firestore:"event_name"`
	Name      string  `firestore:"name"`
	Members   []int64 `firestore:"member_ids"`
	IsOpen    bool    `firestore:"is_open"`
}

func (o Org) ToModel() (*org.Org, error) {
	var members []user.ID
	for i := range o.Members {
		members = append(members, identity.NewID(o.Members[i]))
	}
	id, err := identity.ImportID(o.ID)
	if err != nil {
		return nil, err
	}
	return &org.Org{
		ID: id,
		Event: event.Event{
			ID:   identity.NewID(o.EventID),
			Name: o.EventName,
		},
		Name:    o.Name,
		Members: members,
		IsOpen:  o.IsOpen,
	}, nil
}
