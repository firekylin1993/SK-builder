package service

import (
	"context"

	pb "SK-builder/api/edn/v1"
	"SK-builder/internal/biz"
)

type ReceiverService struct {
	pb.UnimplementedReceiverServer

	uc *biz.ReceiverUsecase
}

func NewReceiverService(uc *biz.ReceiverUsecase) *ReceiverService {
	return &ReceiverService{uc: uc}
}

func (s *ReceiverService) Recrive(ctx context.Context, req *pb.ReceiveRequest) (*pb.ReceiveReply, error) {
	s2, err := s.uc.GetPubKey(ctx, req.Channel)
	if err != nil {
		return nil, err
	}
	return &pb.ReceiveReply{Message: s2}, nil
}
