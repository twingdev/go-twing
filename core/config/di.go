package config

import "github.com/golobby/container/v3"

func init() {

	err := container.Singleton(func() *Config {
		return &Config{}
	})
	if err != nil {
		return
	}
}
