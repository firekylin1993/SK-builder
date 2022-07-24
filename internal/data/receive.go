package data

import (
	"context"

	"SK-builder/internal/biz"
	"SK-builder/internal/infrastructure/db"

	"github.com/go-kratos/kratos/v2/log"
)

type receiverRepo struct {
	data *db.Data
	log  *log.Helper
}

func NewReceiverRepo(data *db.Data, logger log.Logger) biz.ReceiverRepo {
	return &receiverRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *receiverRepo) CreatePubKey(ctx context.Context, channel string) (string, error) {
	return channel, nil
}
