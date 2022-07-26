package service

import (
	pb "SK-builder-demo/api/edn/v1"
	"SK-builder-demo/internal/biz"
	"context"
	"go.opentelemetry.io/otel"
)

type EdnService struct {
	pb.UnimplementedEdnServer

	uc *biz.EdnUsecase
}

func NewEdnService(uc *biz.EdnUsecase) *EdnService {
	return &EdnService{uc: uc}
}

func (s *EdnService) Receiver(ctx context.Context, req *pb.ReceiverRequest) (*pb.ReceiverReply, error) {
	c, span := otel.Tracer("SK-builder-demo.Receiver").Start(ctx, "Receiver")
	defer span.End()
	receiver, err := s.uc.KeyReceive(c, &biz.Edn{
		Channel: req.Channel,
	})

	if err != nil {
		return nil, err
	}

	return &pb.ReceiverReply{Data: receiver.Channel}, nil
}
