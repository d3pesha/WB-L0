package app

import (
	"wb/api"
	"wb/database"
	"wb/repository"
)

type App struct {
	server api.Api
	db     repository.Database
}

func New() App {
	db := database.New()
	return App{
		db:     &db,
		server: api.New(&db),
	}
}

func (a *App) Run() error {
	return a.server.Run()
}
