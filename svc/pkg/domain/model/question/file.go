package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	FileQuestion struct {
		Basic
		FileTypes  FileTypes
		Constraint FileConstraint
	}
	FileTypes struct {
		AcceptAny   bool
		AcceptImage bool
		AcceptPDF   bool
	}
)

const (
	FileQuestionFileTypeField   = "fileTypes"
	FileConstraintsCustomsField = "fileConstraint"
)

func NewFileQuestion(
	id id.QuestionID, text string, fileTypes FileTypes, constraint FileConstraint, formID id.FormID,
) *FileQuestion {
	return &FileQuestion{
		Basic:      NewBasic(id, text, TypeFile, formID),
		FileTypes:  fileTypes,
		Constraint: constraint,
	}
}

func ImportFileQuestion(q StandardQuestion) (*FileQuestion, error) {
	// expect custom field has fileTypes as []bool
	// [AcceptAny, AcceptImage, AcceptPDF]
	fileTypeDataI, has := q.Customs[FileQuestionFileTypeField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for FileQuestion", FileQuestionFileTypeField))
	}
	fileTypeData, ok := fileTypeDataI.([]bool)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be int for FileQuestion", FileQuestionFileTypeField))
	}

	// extend length if not enough
	if len(fileTypeData) < 3 {
		for i := len(fileTypeData); i < 3; i++ {
			fileTypeData = append(fileTypeData, false)
		}
	}

	fileTypes := FileTypes{
		AcceptAny:   fileTypeData[0],
		AcceptImage: fileTypeData[1],
		AcceptPDF:   fileTypeData[2],
	}

	constraintsCustomsData, has := q.Customs[FileConstraintsCustomsField]
	// if FileConstraintsCustomsField is not present, return FileQuestion without constraint
	if !has {
		return NewFileQuestion(q.ID, q.Text, fileTypes, nil, q.FormID), nil
	}

	constraintsCustoms, ok := constraintsCustomsData.(map[string]interface{})
	// if FileConstraintsCustomsField Found, but it is not map[string]interface{}, return error
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileConstraintsCustomsField))
	}

	constraint := NewStandardFileConstraint(fileTypes, constraintsCustoms)
	question := NewFileQuestion(q.ID, q.Text, fileTypes, ImportFileConstraint(constraint), q.FormID)
	return question, nil
}

func (q FileQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})

	qt := []bool{q.FileTypes.AcceptAny, q.FileTypes.AcceptImage, q.FileTypes.AcceptPDF}
	customs[FileQuestionFileTypeField] = qt

	if q.Constraint != nil {
		customs[FileConstraintsCustomsField] = q.Constraint.Export().Customs
	}
	return NewStandardQuestion(TypeFile, q.ID, q.FormID, q.Text, customs)
}
