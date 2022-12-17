package setting

type (
	ThirdParty struct {
		LineLogin LineLogin `yaml:"line_login"`
	}

	LineLogin struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		CallbackURI  string `yaml:"callback_uri"`
		CipherKey    string `yaml:"cipher_key"`
	}
)
