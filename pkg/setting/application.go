package setting

type (
	Application struct {
		Server Server `yaml:"server"`
		Admin  Admin  `yaml:"admin"`
	}
	Server struct {
		Domain          string `yaml:"domain"`
		FrontDomain     string `yaml:"front_domain"`
		FrontDomainPort string `yaml:"front_domain_port"`
	}
	Admin struct {
		JwtSecret string `yaml:"jwt_secret"`
	}
)
