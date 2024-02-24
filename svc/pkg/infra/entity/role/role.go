package entity

const RoleTableName = "Roles"

type Role struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}
