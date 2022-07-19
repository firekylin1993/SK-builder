package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Receiver struct {
	Hello string
}

type ReceiverRepo interface {
	CreatePubKey(ctx context.Context, channel string) (string, error)
}

type ReceiverUsecase struct {
	repo ReceiverRepo
	log  *log.Helper
}

func NewReceiverUsecase(repo ReceiverRepo, logger log.Logger) *ReceiverUsecase {
	return &ReceiverUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ReceiverUsecase) GetPubKey(ctx context.Context, channel string) (string, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %s", channel)
	return uc.repo.CreatePubKey(ctx, channel)
}
