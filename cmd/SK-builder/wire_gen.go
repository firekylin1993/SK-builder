// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"SK-builder-demo/internal/biz"
	"SK-builder-demo/internal/conf"
	"SK-builder-demo/internal/data"
	"SK-builder-demo/internal/data/myotel"
	"SK-builder-demo/internal/data/myrsa"
	"SK-builder-demo/internal/data/mysnowflake"
	"SK-builder-demo/internal/db"
	"SK-builder-demo/internal/server"
	"SK-builder-demo/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireDb(data *conf.Data, logger log.Logger) (*db.Data, func(), error) {
	dbData, cleanup, err := db.NewMysql(data, logger)
	if err != nil {
		return nil, nil, err
	}
	return dbData, func() {
		cleanup()
	}, nil
}

// wireApp init kratos application.
func wireApp(confServer *conf.Server, dbData *db.Data, logger log.Logger) (*kratos.App, error) {
	ednRepo := data.NewEdnRepo(dbData, logger)
	ednUsecase := biz.NewEdnUsecase(ednRepo, logger)
	ednService := service.NewEdnService(ednUsecase)
	grpcServer := server.NewGRPCServer(confServer, ednService, logger)
	httpServer := server.NewHTTPServer(confServer, ednService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, nil
}

func wireOtel(contextContext context.Context, confServer *conf.Server, logger log.Logger) (func(), error) {
	client := myotel.NewMetricClient(confServer)
	exporter := myotel.NewMetricExporter(contextContext, client)
	otlptraceClient := myotel.NewTracerClient(confServer)
	otlptraceExporter := myotel.NewTracerExporter(contextContext, otlptraceClient)
	v, err := newOtel(contextContext, confServer, exporter, otlptraceExporter)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func wireBucket(contextContext context.Context, confServer *conf.Server, dbData *db.Data, logger log.Logger) error {
	privateKey := myrsa.NewProviderKey()
	publicKey := myrsa.NewPublicKey()
	rsaKey := myrsa.NewRsaKey(confServer, privateKey, publicKey)
	snowNode := mysnowflake.NewSnowNode(confServer)
	rsaBucketRepo := myrsa.NewBucketRepo(dbData, logger)
	rsaBucket := myrsa.NewRsaBucket(confServer, rsaKey, snowNode, rsaBucketRepo)
	error2 := newBucket(contextContext, rsaBucket, logger)
	return error2
}
