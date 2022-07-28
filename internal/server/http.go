package server

import (
	v1 "SK-builder-demo/api/edn/v1"
	"SK-builder-demo/internal/conf"
	"SK-builder-demo/internal/data/p8s"
	"SK-builder-demo/internal/service"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	edn *service.EdnService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(p8s.MetricSeconds)),
				metrics.WithRequests(prom.NewCounter(p8s.MetricRequests)),
			),
			validate.Validator(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	srv.Handle("/metrics", promhttp.Handler())

	h := openapiv2.NewHandler()
	// swagger
	srv.HandlePrefix("/q/", h)

	v1.RegisterEdnHTTPServer(srv, edn)
	return srv
}
