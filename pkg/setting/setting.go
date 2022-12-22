package setting

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envLocation := os.Getenv("ENV_LOCATION")
	log.Println("ENV_LOCATION: " + envLocation)
	reader, err := os.Open(envLocation)
	if err != nil {
		dir, _ := os.Getwd()
		log.Fatalln(nil, "failed to open setting file: %v, %v\n", dir, err)
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
