package rpc

import (
	"context"

	"go.kicksware.com/api/service-common/core"
	"google.golang.org/grpc"

	"github.com/timoth-y/scrapnote-api/data.records/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"github.com/timoth-y/scrapnote-api/data.records/core/repo"
)

//go:generate protoc --proto_path=proto/ --go_out=plugins=grpc,paths=source_relative:proto/. record.proto

type Handler struct {
	repo repo.RecordRepository
	auth core.AuthService
}

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterRecordServiceServer(server, handler)
	}
}

func NewHandler(repo repo.RecordRepository) *Handler {
	return &Handler{
		repo,
		nil,
	}
}

func (h Handler) Get(ctx context.Context, filter *proto.RecordFilter) (r *proto.RecordResponse, err error) {
	var records []*model.Record

	if len(filter.RecordID) > 0 {
		records, err = h.repo.Retrieve(filter.RecordID)
	} else if len(filter.TopicID) > 0 {
		records, err = h.repo.RetrieveBy(filter.TopicID)
	} else {
		records, err = h.repo.RetrieveAll()
	}

	r = &proto.RecordResponse{
		Records: proto.NativeToRecords(records),
		Count: int64(len(records)),
	}
	return
}

func (h Handler) Count(ctx context.Context, filter *proto.RecordFilter) (*proto.RecordResponse, error) {
	panic("implement me")
}
