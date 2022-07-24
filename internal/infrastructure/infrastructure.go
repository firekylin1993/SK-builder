package infrastructure

import (
	"SK-builder/internal/infrastructure/db"
	"SK-builder/internal/infrastructure/mykey"
	"SK-builder/internal/infrastructure/myotel"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewDB,

	myotel.NewTracerClient,
	myotel.NewTracerExporter,
	myotel.NewMetricClient,
	myotel.NewMetricExporter,

	mykey.NewPublicKey,
	mykey.NewProviderKey,
	mykey.NewRsaKey,

	mykey.NewSnowNode,

	mykey.NewRsaBucket,
)
