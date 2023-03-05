package relation

const (
	RelationRootName    = "Relation"
	RelationOrgUserName = "OrgUser"
)

type OrgUserRelation struct {
	OrgID  string `json:"org_id"`
	UserID string `json:"user_id"`
}
