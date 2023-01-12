package setting

type (
	Application struct {
		Server Server `yaml:"server"`
		Admin  Admin  `yaml:"admin"`
	}
	Server struct {
		OnProduction bool     `yaml:"on_production"`
		Frontend     Frontend `yaml:"frontend"`
		Backend      Backend  `yaml:"backend"`
	}
	Admin struct {
		JwtSecret string `yaml:"jwt_secret"`
	}
	Frontend struct {
		Protocol string `yaml:"protocol"`
		Domain   string `yaml:"domain"`
		Port     string `yaml:"port"`
	}
	Backend struct {
		Protocol string `yaml:"protocol"`
		Domain   string `yaml:"domain"`
		Port     string `yaml:"port"`
	}
)
