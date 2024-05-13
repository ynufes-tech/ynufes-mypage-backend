package question

import "fmt"

type (
	FileType               int
	StandardFileConstraint struct {
		Type    FileType
		Customs map[string]interface{}
	}
	FileConstraint interface {
		GetFileType() FileType
		GetExtensions() []Extension
		Export() (*StandardFileConstraint, error)
		ValidateFiles(file []File) error
	}
	File struct {
		FileName string
		Data     []byte
	}
	Extension string
)

const (
	Image               FileType = 1
	PDF                 FileType = 2
	FileTypeCustomField          = "type"
)

func NewStandardFileConstraint(fileType FileType, customs map[string]interface{}) (*StandardFileConstraint, error) {
	customs[FileTypeCustomField] = fileType
	return &StandardFileConstraint{
		Type:    fileType,
		Customs: customs,
	}, nil
}

func ImportFileConstraint(st StandardFileConstraint) (FileType, FileConstraint, error) {
	switch st.Type {
	case Image:
		return Image, ImportImageFileConstraint(st), nil
	default:
		return 0, nil, fmt.Errorf("invalid file type")
	}
}
