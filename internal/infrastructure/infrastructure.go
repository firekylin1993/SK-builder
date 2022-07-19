package infrastructure

import (
	"SK-builder/internal/infrastructure/otel/trace"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(trace.NewTracerClient, trace.NewTracerExporter)
