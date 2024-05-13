package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	FileQuestion struct {
		Basic
		FileTypes   FileTypes
		Constraints map[FileType]FileConstraint
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
	id id.QuestionID, text string, fileTypes FileTypes, constraint map[FileType]FileConstraint, formID id.FormID,
) *FileQuestion {
	return &FileQuestion{
		Basic:       NewBasic(id, text, TypeFile, formID),
		FileTypes:   fileTypes,
		Constraints: constraint,
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
	//if FileConstraintsCustomsField is not present, return FileQuestion without constraint
	if !has {
		return NewFileQuestion(q.ID, q.Text, fileTypes, nil, q.FormID), nil
	}

	constraintsCustoms, ok := constraintsCustomsData.([]interface{})
	//if FileConstraintsCustomsField Found, but it is not slice, return error
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileConstraintsCustomsField))
	}

	constraints := make(map[FileType]FileConstraint, len(constraintsCustoms))

	for _, constraintCustomsDataI := range constraintsCustoms {
		constraintCustomsData, ok := constraintCustomsDataI.(map[string]interface{})
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileConstraintsCustomsField))
		}

		st, err := NewStandardFileConstraint(FileType(constraintCustomsData[FileTypeCustomField].(float64)), constraintCustomsData)
		if err != nil {
			return nil, fmt.Errorf("failed to import file constraint: %w", err)
		}
		fileType, constraint, err := ImportFileConstraint(*st)
		if err != nil {
			return nil, fmt.Errorf("failed to import file constraint: %w", err)
		}
		constraints[fileType] = constraint
	}

	question := NewFileQuestion(q.ID, q.Text, fileTypes, constraints, q.FormID)
	return question, nil
}

func (q FileQuestion) Export() (*StandardQuestion, error) {
	customs := make(map[string]interface{})

	qt := []bool{q.FileTypes.AcceptAny, q.FileTypes.AcceptImage, q.FileTypes.AcceptPDF}
	customs[FileQuestionFileTypeField] = qt

	if q.Constraints != nil {
		constraints := make([]map[string]interface{}, 0, len(q.Constraints))
		for _, constraint := range q.Constraints {
			c, err := constraint.Export()
			if err != nil {
				return nil, fmt.Errorf("failed to export file constraint: %w", err)
			}
			constraints = append(constraints, c.Customs)
		}
	}
	return NewStandardQuestion(TypeFile, q.ID, q.FormID, q.Text, customs), nil
}
