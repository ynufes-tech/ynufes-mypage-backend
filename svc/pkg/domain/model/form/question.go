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
	Question struct {
		ID           QID
		Type         QuestionType
		QuestionText string
		Order        int
		Properties   QuestionProperties
	}
	QuestionType       int
	QID                util.ID
	QuestionProperties interface {
		Export() map[string]interface{}
	}
)
