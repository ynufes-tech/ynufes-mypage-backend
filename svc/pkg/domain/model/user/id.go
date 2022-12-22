package user

import "strconv"

type ID int64

func NewID(id int64) ID {
	return ID(id)
}

func ImportID(id string) (ID, error) {
	result, err := strconv.ParseInt(id, 36, 64)
	if err != nil {
		return 0, err
	}
	return ID(result), nil
}

func (i ID) ExportID() string {
	return strconv.FormatInt(int64(i), 36)
}
