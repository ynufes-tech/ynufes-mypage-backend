package form

import "ynufes-mypage-backend/svc/pkg/domain/model/util"

const (
	CheckBox    QuestionType = 1
	RadioButton QuestionType = 2
	TextField   QuestionType = 3
	TextArea    QuestionType = 4
	File        QuestionType = 5
)

type (
	Question interface {
		ID() QID
		Type() QuestionType
		Text() string
		Order() int
	}
	QuestionType int
	QID          util.ID

	QuestionBasic struct {
		QID           QID
		QuestionText  string
		QuestionOrder int
	}
)

func (q QuestionBasic) ID() QID {
	return q.QID
}

func (q QuestionBasic) Order() int {
	return q.QuestionOrder
}

func (q QuestionBasic) Text() string {
	return q.QuestionText
}

type (
	CheckboxQuestion struct {
		QuestionBasic
		QID     QID
		Options []CheckboxOption
	}
	CheckboxOption struct {
		ID    util.ID
		Text  string
		Order int
	}
)

func (q CheckboxQuestion) Type() QuestionType {
	return CheckBox
}
