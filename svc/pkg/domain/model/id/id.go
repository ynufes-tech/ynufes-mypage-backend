package id

type ID interface {
	ExportID() string
	HasValue() bool
}
