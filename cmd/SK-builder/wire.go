//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"SK-Builder/internal/biz"
	"SK-Builder/internal/conf"
	"SK-Builder/internal/data"
	"SK-Builder/internal/db"
	"SK-Builder/internal/server"
	"SK-Builder/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireDb(*conf.Data, log.Logger) (*db.Data, func(), error) {
	panic(
		wire.Build(
			db.ProviderSet,
		),
	)
}

// wireApp init kratos application.
func wireApp(*conf.Server, *db.Data, log.Logger) (*kratos.App, error) {
	panic(
		wire.Build(
			server.ProviderSet,
			data.ProviderSet,
			biz.ProviderSet,
			service.ProviderSet,
			newApp,
		),
	)
}

func wireOtel(context.Context, *conf.Server, log.Logger) (func(), error) {
	panic(
		wire.Build(
			data.ProviderSet,
			newOtel,
		),
	)
}

func wireBucket(context.Context, *conf.Server, *db.Data, log.Logger) error {
	panic(
		wire.Build(
			data.ProviderSet,
			newBucket,
		),
	)
}
