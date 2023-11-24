package repositoryimpl

import (
	"backend-demo/ent"
	"backend-demo/ent/event"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
	"time"
)

type eventRepoImpl struct {
	DBClient *database.Client
}

// CreateEvent implements repository.EventRepository.
func (e *eventRepoImpl) CreateEvent(c context.Context, name string, userId int) (*ent.Event, error) {
	newEvent, err := e.DBClient.Client.Event.Create().SetName(name).SetUsersID(userId).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// DeleteEvent implements repository.EventRepository.
func (e *eventRepoImpl) DeleteEvent(c context.Context, eventId int) (*ent.Event, error) {
	deleteEvent, err := e.DBClient.Client.Event.UpdateOneID(eventId).SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return deleteEvent, nil
}

// GetEvents implements repository.EventRepository.
func (e *eventRepoImpl) GetEvents(c context.Context) ([]*ent.Event, error) {
	events, err := e.DBClient.Client.Event.Query().Where(event.DeletedAtIsNil()).All(context.Background())
	if err != nil {
		return nil, err
	}

	return events, nil
}

// GetEventsById implements repository.EventRepository.
func (e *eventRepoImpl) GetEventsById(c context.Context, userId int) ([]*ent.Event, error) {
	events, err := e.DBClient.Client.Event.Query().Where(event.And(event.CreatedByEQ(userId), event.DeletedAtIsNil())).All(context.Background())
	if err != nil {
		return nil, err
	}

	return events, err
}

// UpdateEvent implements repository.EventRepository.
func (e *eventRepoImpl) UpdateEvent(c context.Context, eventId int) (*ent.Event, error) {
	updateEvent, err := e.DBClient.Client.Event.UpdateOneID(eventId).Where(event.DeletedAtIsNil()).SetName("updateEvent").Save(context.Background())
	if err != nil {
		return nil, err
	}

	return updateEvent, nil
}

func NewEventRepoImpl(client *database.Client) repository.EventRepository {
	return &eventRepoImpl{
		DBClient: client,
	}
}
