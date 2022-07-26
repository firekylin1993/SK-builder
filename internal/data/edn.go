package data

import (
	"SK-Builder/internal/biz"
	"SK-Builder/internal/db"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ednRepo struct {
	data *db.Data
	log  *log.Helper
}

func NewEdnRepo(data *db.Data, logger log.Logger) biz.EdnRepo {
	return &ednRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ednRepo) Receive(ctx context.Context, g *biz.Edn) (*biz.Edn, error) {
	return g, nil
}
