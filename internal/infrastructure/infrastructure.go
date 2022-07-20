package infrastructure

import (
	"SK-builder/internal/infrastructure/myotel"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	myotel.NewTracerClient, 
	myotel.NewTracerExporter,
	myotel.NewMetricClient,
	myotel.NewMetricExporter,
)
