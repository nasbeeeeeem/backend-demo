package repositoryimpl

import (
	"backend-demo/pkg/domain/model"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
)

type eventRepoImpl struct {
	DBClient *database.Engine
}

// CreateEvent implements repository.EventRepository.
func (*eventRepoImpl) CreateEvent(c context.Context, name string, userId int) (*model.Event, error) {
	panic("unimplemented")
}

// DeleteEvent implements repository.EventRepository.
func (*eventRepoImpl) DeleteEvent(c context.Context, eventId int) (*model.Event, error) {
	panic("unimplemented")
}

// GetEvents implements repository.EventRepository.
func (*eventRepoImpl) GetEvents(c context.Context) ([]*model.Event, error) {
	panic("unimplemented")
}

// GetEventsById implements repository.EventRepository.
func (*eventRepoImpl) GetEventsById(c context.Context, userId int) ([]*model.Event, error) {
	panic("unimplemented")
}

// UpdateEvent implements repository.EventRepository.
func (*eventRepoImpl) UpdateEvent(c context.Context, eventId int) (*model.Event, error) {
	panic("unimplemented")
}

func NewEventRepoImpl(client *database.Engine) repository.EventRepository {
	return &eventRepoImpl{
		DBClient: client,
	}
}
