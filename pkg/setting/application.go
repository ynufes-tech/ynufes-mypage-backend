package setting

type (
	Application struct {
		Server Server `yaml:"server"`
	}
	Server struct {
		Domain string `yaml:"domain"`
	}
)
