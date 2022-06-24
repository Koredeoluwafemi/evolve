//go:build darwin
// +build darwin

package config

var App struct {
	Name   string `env:"Name" envDefault:"EvolveData"`
	Mode   string `env:"Mode" envDefault:"test"`
	Port   string `env:"Port" envDefault:"4555"`
	ENV    string `env:"ENV" envDefault:"local"`
	JWTKey string `env:"JWTKey" envDefault:"kiuru72h2ywn"`
}
