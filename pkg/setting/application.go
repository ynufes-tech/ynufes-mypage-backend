package setting

type (
	Application struct {
		Server Server `yaml:"server"`
		Admin  Admin  `yaml:"admin"`
	}
	Server struct {
		Domain string `yaml:"domain"`
	}
	Admin struct {
		JwtSecret string `yaml:"jwt_secret"`
	}
)
