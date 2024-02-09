// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/mateusmacedo/go-nunu-template/internal/server"
	"github.com/mateusmacedo/go-nunu-template/pkg/app"
	"github.com/mateusmacedo/go-nunu-template/pkg/log"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	task := server.NewTask(logger)
	appApp := newApp(task)
	return appApp, func() {
	}, nil
}

// wire.go:

var taskSet = wire.NewSet(server.NewTask)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(app.WithServer(task), app.WithName("demo-task"))
}