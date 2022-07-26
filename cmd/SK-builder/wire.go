//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"SK-builder-demo/internal/biz"
	"SK-builder-demo/internal/conf"
	"SK-builder-demo/internal/data"
	"SK-builder-demo/internal/db"
	"SK-builder-demo/internal/server"
	"SK-builder-demo/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
