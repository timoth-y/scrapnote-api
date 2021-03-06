package business

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"github.com/timoth-y/scrapnote-api/lib.common/api/events"
	"github.com/timoth-y/scrapnote-api/lib.common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/core/service"
)

type recordService struct {
	events *events.Broker
	config config.ServiceConfig
}

func NewRecordService(config config.ServiceConfig, serializer core.Serializer) service.RecordService {
	return &recordService {
		events.NewEventsBroker(config.Events, "amq.topic", serializer),
		config,
	}
}

func (s *recordService) GetOne(ctx context.Context, id string) (*model.Record, error) {
	panic("implement me")
}

func (s *recordService) Get(ctx context.Context, topic string) ([]*model.Record, error) {
	panic("implement me")
}

func (s *recordService) Add(ctx context.Context, record *model.Record) error {
	return s.events.Emmit(ctx,"records.add", record)
}

func (s *recordService) Update(ctx context.Context, record *model.Record) error {
	return s.events.Emmit(ctx,"records.update", record)
}

func (s *recordService) Delete(ctx context.Context, id string) error {
	return s.events.Emmit(ctx, "records.delete", id)
}