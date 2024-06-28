package app

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/slavaromanov/wire-example/postgres"
)

type Config struct {
	RepoDSN postgres.DSN `envconfig:"psql_dsn" default:"postgres://postgres:postgres@localhost/postgres?sslmode=disable"`
}

func NewConfig() *Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return &cfg
}
