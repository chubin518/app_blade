//go:build wireinject
// +build wireinject

package cmd

import (
	"app_blade/internal/controller"
	"app_blade/internal/repository"
	"app_blade/internal/service"
	"app_blade/internal/startup"
	"app_blade/pkg/app"
	"app_blade/pkg/config"
	"app_blade/pkg/database"
	"app_blade/pkg/logging"
	"app_blade/pkg/web"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	logging.ProviderSet,
	web.ProviderSet,
	app.ProviderSet,
	startup.ProviderSet,
	controller.ProviderSet,
	service.ProviderSet,
	repository.ProviderSet,
	database.ProviderSet,
)

func CreateApp(options ...config.Option) (*app.App, error) {
	panic(wire.Build(providerSet))
}
