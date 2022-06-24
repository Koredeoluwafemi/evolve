package config

import "github.com/caarlos0/env/v6"

var Environment struct {
	Local string `env:"Local" envDefault:"local"`
	Stage string `env:"Stage" envDefault:"stage"`
}

func init() {
	_ = env.Parse(&App)
	_ = env.Parse(&Environment)
}
