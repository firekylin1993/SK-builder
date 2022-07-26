package data

import (
	"SK-Builder/internal/data/myotel"
	"SK-Builder/internal/data/myrsa"
	"SK-Builder/internal/data/mysnowflake"
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
