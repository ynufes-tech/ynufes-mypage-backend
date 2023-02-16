package setting

type (
	Service struct {
		Authentication Authentication `yaml:"authentication"`
	}
	Authentication struct {
		SecureCookie bool   `yaml:"secure_cookie"`
		JwtSecret    string `yaml:"jwt_secret"`
	}
)
