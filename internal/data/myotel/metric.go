package myotel

import (
	"SK-builder-demo/internal/conf"
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"google.golang.org/grpc"

	"go.opentelemetry.io/otel/sdk/metric/export"
)

func NewMetricClient(c *conf.Server) otlpmetric.Client {
	if c == nil || c.Otel == nil || c.Otel.Addr == "" {
		return nil
	}
	return otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(c.Otel.Addr),
		otlpmetricgrpc.WithDialOption(grpc.WithBlock()))
}

func NewMetricExporter(ctx context.Context, client otlpmetric.Client) export.Exporter {
	if client == nil {
		return nil
	}
	metricExp, err := otlpmetric.New(ctx, client)
	if err != nil {
		panic(errors.WithMessage(err, "初始化 metric exporter 失败"))
	}
	return metricExp
}
