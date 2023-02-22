package question

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
)

type Basic struct {
	ID      ID
	Text    string
	EventID event.ID
	FormID  form.ID
	qType   Type
}

func NewBasic(id ID, text string, eventID event.ID, formID form.ID, qType Type) Basic {
	return Basic{
		ID:      id,
		Text:    text,
		EventID: eventID,
		FormID:  formID,
		qType:   qType,
	}
}

func (b Basic) GetID() ID {
	return b.ID
}

func (b Basic) GetText() string {
	return b.Text
}

func (b Basic) GetEventID() event.ID {
	return b.EventID
}

func (b Basic) GetFormID() form.ID {
	return b.FormID
}

func (b Basic) GetType() Type {
	return b.qType
}
