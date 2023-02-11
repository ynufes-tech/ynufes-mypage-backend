package entity

import "ynufes-mypage-backend/svc/pkg/domain/model/event"

const EventCollectionName = "Events"

type Event struct {
	ID   event.ID `firestore:"-"`
	Name string   `firestore:"name"`
}

func (e Event) ToModel() (*event.Event, error) {
	return &event.Event{
		ID:   e.ID,
		Name: e.Name,
	}, nil
}
