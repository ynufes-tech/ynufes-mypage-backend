package entity

const StaffTableName = "Staffs"

type Staff struct {
	UserID  string `json:"-"`
	IsAdmin bool   `json:"is_admin"`
}
