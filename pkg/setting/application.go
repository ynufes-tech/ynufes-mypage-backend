package setting

type (
	Application struct {
		Server Server `yaml:"server"`
		Admin  Admin  `yaml:"admin"`
	}
	Server struct {
		Domain         string `yaml:"domain"`
		FrontendDomain string `yaml:"frontend_domain"`
	}
	Admin struct {
		JwtSecret string `yaml:"jwt_secret"`
	}
)
