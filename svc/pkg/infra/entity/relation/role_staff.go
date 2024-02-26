package entity

const RelationRoleStaffName = "RoleStaff"

type RoleStaffRelation struct {
	RoleID string `json:"role_id"`
	UserID string `json:"user_id"`
}
