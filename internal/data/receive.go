package data

import (
	"context"

	"SK-builder/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type receiverRepo struct {
	data *Data
	log  *log.Helper
}

func NewReceiverRepo(data *Data, logger log.Logger) biz.ReceiverRepo {
	return &receiverRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *receiverRepo) CreatePubKey(ctx context.Context, channel string) (string, error) {
	return "qqq", nil
}