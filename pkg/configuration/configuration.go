package configuration

import (
	env "github.com/Netflix/go-env"
)

type Config struct {
	HTTP_SOCKET  string `env:"HTTP_SOCKET"`
	POSTGRES_URL string `env:"POSTGRES_URL"`
}

func Load() (*Config, error) {
	var output Config
	_, err := env.UnmarshalFromEnviron(&output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
