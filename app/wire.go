//go:build wireinject

package app

import (
	"context"
	"github.com/google/wire"
	"github.com/slavaromanov/wire-example/postgres"
)

func NewApp(ctx context.Context) (*App, error) {
	wire.Build(wire.NewSet(
		newApp,
		NewConfig,
		wire.FieldsOf(new(*Config), "RepoDSN"),
		postgres.NewRepo,
		wire.Bind(new(Repo), new(*postgres.Repo)),
	))
	return &App{}, nil
}
