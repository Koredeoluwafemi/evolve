//go:build linux
// +build linux

package config

var App struct {
	Name    string `env:"Name" envDefault:"EvolveData"`
	Url     string `env:"Url" envDefault:""`
	BaseUrl string `env:"BaseUrl" envDefault:""`
	JWTKey  string `env:"JWTKey" envDefault:"SGa37bXXtT1ZfkB1maTha3h9jLJQpEpd-dZ7aYqEvkB5M"`
	Mode    string `env:"Mode"   envDefault:"test"`
	Port    string `env:"Port" envDefault:"3500"`
	ENV     string `env:"ENV" envDefault:"stage"`
}
