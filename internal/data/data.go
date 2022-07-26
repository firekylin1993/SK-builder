package data

import (
	"SK-builder-demo/internal/data/myotel"
	"SK-builder-demo/internal/data/myrsa"
	"SK-builder-demo/internal/data/mysnowflake"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewGreeterRepo,
	NewEdnRepo,

	mysnowflake.NewSnowNode,

	myrsa.NewBucketRepo,
	myrsa.NewRsaBucket,
	myrsa.NewProviderKey,
	myrsa.NewPublicKey,
	myrsa.NewRsaKey,

	myotel.NewTracerClient,
	myotel.NewTracerExporter,
	myotel.NewMetricClient,
	myotel.NewMetricExporter,
)
