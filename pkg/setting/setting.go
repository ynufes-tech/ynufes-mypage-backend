package setting

var (
	Setting setting
)

type (
	setting struct {
		Application    Application    `yaml:"application"`
		Infrastructure Infrastructure `yaml:"infrastructure"`
		ThirdParty     ThirdParty     `yaml:"third_party"`
		Service        Service        `yaml:"service"`
	}
)
