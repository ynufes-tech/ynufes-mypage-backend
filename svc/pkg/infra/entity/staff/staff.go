package entity

const StaffTableName = "Staffs"

type Staff struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
}
