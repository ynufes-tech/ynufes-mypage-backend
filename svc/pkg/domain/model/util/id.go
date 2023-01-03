package util

import (
	"ynufes-mypage-backend/pkg/identity"
)

type ID interface {
	ExportID() string
	HasValue() bool
}

func ImportID(id string) (ID, error) {
	return identity.ImportID(id)
}
