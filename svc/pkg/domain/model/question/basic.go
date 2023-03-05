package question

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
)

type Basic struct {
	ID        id.QuestionID
	Text      string
	EventID   id.EventID
	FormID    id.FormID
	SectionID id.SectionID
	qType     Type
}

func NewBasic(id id.QuestionID, text string, eventID id.EventID, qType Type) Basic {
	return Basic{
		ID:      id,
		Text:    text,
		EventID: eventID,
		qType:   qType,
	}
}

func (b Basic) GetID() id.QuestionID {
	return b.ID
}

func (b Basic) GetText() string {
	return b.Text
}

func (b Basic) GetEventID() id.EventID {
	return b.EventID
}

func (b Basic) GetFormID() id.FormID {
	return b.FormID
}

func (b Basic) GetType() Type {
	return b.qType
}

func (b *Basic) AssignID(id id.QuestionID) error {
	if b.ID != nil && b.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	b.ID = id
	return nil
}
