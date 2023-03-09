package question

type (
	StandardFileConstraint struct {
		Type    FileType
		Customs map[string]interface{}
	}
	FileConstraint interface {
		GetFileType() FileType
		GetExtensions() []Extension
		Export() StandardFileConstraint
		ValidateFiles(file []File) error
	}
	File struct {
		FileName string
		Data     []byte
	}
	Extension string
)

func NewStandardFileConstraint(fileType FileType, customs map[string]interface{}) StandardFileConstraint {
	return StandardFileConstraint{
		Type:    fileType,
		Customs: customs,
	}
}

func ImportFileConstraint(standard StandardFileConstraint) FileConstraint {
	switch standard.Type {
	case Image:
		return ImportImageFileConstraint(standard)
	default:
		return nil
	}
}
