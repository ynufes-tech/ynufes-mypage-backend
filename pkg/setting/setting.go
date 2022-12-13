package setting

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	setting *Setting = nil
)

type (
	Setting struct {
		Application    Application    `yaml:"application"`
		Infrastructure Infrastructure `yaml:"infrastructure"`
		ThirdParty     ThirdParty     `yaml:"third_party"`
		Service        Service        `yaml:"service"`
	}
)

func init() {
	reader, err := os.Open(os.Getenv("ENV_LOCATION"))
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(reader)
	setting = &Setting{}
	err = decoder.Decode(setting)
	if err != nil {
		panic(err)
	}
}

func Get() Setting {
	if setting == nil {
		panic("setting is nil")
	}
	return *setting
}
