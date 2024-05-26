package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	FileQuestion struct {
		Basic
		FileTypes FileTypes
		ImageFileConstraint
	}
	FileTypes struct {
		AcceptAny   bool
		AcceptImage bool
		AcceptPDF   bool
	}
)

const (
	FileQuestionFileTypeField = "fileTypes"
	FileImageConstraintField  = "img_c"
)

func NewFileQuestion(
	id id.QuestionID, text string, fileTypes FileTypes,
	imgConstraint ImageFileConstraint,
	formID id.FormID,
) *FileQuestion {
	return &FileQuestion{
		Basic:               NewBasic(id, text, TypeFile, formID),
		FileTypes:           fileTypes,
		ImageFileConstraint: imgConstraint,
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

	imgConstraintCustomR, has := q.Customs[FileImageConstraintField]
	//if FileConstraintsCustomsField is not present, return FileQuestion without constraint
	if !has {
		return NewFileQuestion(q.ID, q.Text, fileTypes, ImageFileConstraint{}, q.FormID), nil
	}

	imgConstraintCustom, ok := imgConstraintCustomR.(map[string]interface{})
	//if FileConstraintsCustomsField Found, but it is not slice, return error
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileImageConstraintField))
	}
	imgConstraint, err := ImportImageFileConstraint(imgConstraintCustom)
	if err != nil {
		return nil, fmt.Errorf("failed to import ImageFileConstraint: %w", err)
	}
	question := NewFileQuestion(q.ID, q.Text, fileTypes, *imgConstraint, q.FormID)
	return question, nil
}

func (q FileQuestion) Export() (*StandardQuestion, error) {
	customs := make(map[string]interface{})

	qt := []bool{q.FileTypes.AcceptAny, q.FileTypes.AcceptImage, q.FileTypes.AcceptPDF}
	customs[FileQuestionFileTypeField] = qt
	customs[FileImageConstraintField] = q.ImageFileConstraint.Export()
	return NewStandardQuestion(TypeFile, q.ID, q.FormID, q.Text, customs), nil
}
