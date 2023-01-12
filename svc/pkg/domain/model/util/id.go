package util

type ID interface {
	ExportID() string
	HasValue() bool
}
