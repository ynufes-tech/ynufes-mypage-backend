package setting

type (
	Infrastructure struct {
		Firestore Firestore `yaml:"firestore"`
		Firebase  Firebase  `yaml:"firebase"`
	}
	Firestore struct {
		ProjectID          string `yaml:"project_id"`
		JsonCredentialFile string `yaml:"json_credential_file"`
	}
	Firebase struct {
		DatabaseURL        string `yaml:"database_url"`
		JsonCredentialFile string `yaml:"json_credential_file"`
	}
)
