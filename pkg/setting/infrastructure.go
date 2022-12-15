package setting

type (
	Infrastructure struct {
		Firestore Firestore `yaml:"firestore"`
	}
	Firestore struct {
		ProjectID          string `yaml:"project_id"`
		JsonCredentialFile string `yaml:"json_credential_file"`
	}
)
