package trace

import (
	"SK-builder/internal/conf"
	"context"
	"errors"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
)

func NewTracerClient(c *conf.Server) otlptrace.Client {
	if c == nil || c.Otel == nil || c.Otel.Addr == "" {
		return nil
	}
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(c.Otel.Addr),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	return traceClient
}

func NewTracerExporter(ctx context.Context, client otlptrace.Client) *otlptrace.Exporter {
	if client == nil {
		return nil
	}
	traceExp, err := otlptrace.New(ctx, client)
	if err != nil {
		panic(errors.New("failed to create trace exporter"))
	}
	return traceExp
}
