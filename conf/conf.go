package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Default params for our application.
const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

// Vars respresent configuration for an app.
type Vars struct {
	AppName      string `toml:"-"`
	Mode         string `toml:"mode"`
	Port         string `toml:"port"`
	GraceTimeout int    `toml:"grace_timeout"`
	Log          struct {
		Access string `toml:"access"`
		Error  string `toml:"error"`
		Format string `toml:"format"`
	} `toml:"log"`
	Postgres struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		DB   string `toml:"db"`
	} `toml:"postgres"`
}

//Read application configuration into Vars struct.
func Read(appName, path string) *Vars {
	var conf = Vars{AppName: appName}
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}
