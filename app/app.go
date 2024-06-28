package app

import "context"

type Repo interface {
	Ping(ctx context.Context) error
}

type App struct {
	repo Repo
}

func newApp(repo Repo) *App {
	return &App{repo: repo}
}

func (a *App) Run(ctx context.Context) error {
	return a.repo.Ping(ctx)
}
