//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"SK-builder/internal/biz"
	"SK-builder/internal/conf"
	"SK-builder/internal/data"
	"SK-builder/internal/infrastructure"
	"SK-builder/internal/server"
	"SK-builder/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(context.Context, *conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, infrastructure.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newBucket, newApp))
}

func wireProvider(context.Context, *conf.Server, log.Logger) (func(), error) {
	panic(wire.Build(infrastructure.ProviderSet, newProvider))
}

// func wireBucket(context.Context, *conf.Data, *conf.Server, log.Logger) error {
// 	panic(wire.Build(data.ProviderSet, infrastructure.ProviderSet, NewBucket))
// }
