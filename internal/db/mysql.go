package db

import (
	"SK-Builder/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

// NewMysql .
func NewMysql(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
