package setting

type (
	Application struct {
		Server         Server         `yaml:"server"`
		Authentication Authentication `yaml:"authentication"`
	}
	Server struct {
		OnProduction bool     `yaml:"on_production"`
		Frontend     Frontend `yaml:"frontend"`
		Backend      Backend  `yaml:"backend"`
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
