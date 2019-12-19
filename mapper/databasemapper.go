package mapper

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Databases struct {
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Root     string `toml:"root"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type Application struct {
	Databases *Databases
}

var application = new(Application)

func AnalysisApplication() *Application {
	_, err := toml.DecodeFile("application.toml", application)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	return application
}
