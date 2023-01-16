package util

type ID interface {
	ExportID() string
	HasValue() bool
	GetValue() int64
}
