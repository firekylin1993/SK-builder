package biz

import (
	"context"

	v1 "SK-builder/api/edn/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrChannelNotFound = errors.NotFound(v1.ErrorReason_CHANNEL_NOT_FOUND.String(), "user channel found")
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
