//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/mateusmacedo/go-notification/internal/repository"
	"github.com/mateusmacedo/go-notification/internal/server"
	"github.com/mateusmacedo/go-notification/pkg/app"
	"github.com/mateusmacedo/go-notification/pkg/log"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
)

// build App
func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		server.NewMigrate,
		newApp,
	))
}
