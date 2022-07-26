package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Edn struct {
	Channel string
}

type EdnRepo interface {
	Receive(context.Context, *Edn) (*Edn, error)
}

type EdnUsecase struct {
	repo EdnRepo
	log  *log.Helper
}

func NewEdnUsecase(repo EdnRepo, logger log.Logger) *EdnUsecase {
	return &EdnUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *EdnUsecase) KeyReceive(ctx context.Context, g *Edn) (*Edn, error) {
	uc.log.WithContext(ctx).Infof("KeyReceive: %v", g.Channel)
	return uc.repo.Receive(ctx, g)
}
