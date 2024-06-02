package question

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
