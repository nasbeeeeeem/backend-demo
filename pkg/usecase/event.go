package usecase

import (
	"backend-demo/pkg/domain/model"
	"backend-demo/pkg/domain/repository"
	"context"
	"time"
)

type EventUseCase interface {
	Create(c context.Context, name string, userId int) (*model.Event, error)
	Events(c context.Context) ([]*model.Event, error)
	Update(c context.Context, eventId int) (*model.Event, error)
	Delete(c context.Context, eventId int) (*model.Event, error)
}

type eventUseCase struct {
	repository repository.EventRepository
	timeout    time.Duration
}

// Create implements EventUseCase.
func (eu *eventUseCase) Create(c context.Context, name string, userId int) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(c, eu.timeout)
	defer cancel()

	newEvent, err := eu.repository.CreateEvent(ctx, name, userId)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

// Events implements EventUseCase.
func (eu *eventUseCase) Events(c context.Context) ([]*model.Event, error) {
	ctx, cancel := context.WithTimeout(c, eu.timeout)
	defer cancel()

	events, err := eu.repository.GetEvents(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Update implements EventUseCase.
func (eu *eventUseCase) Update(c context.Context, eventId int) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(c, eu.timeout)
	defer cancel()

	updateEvent, err := eu.repository.UpdateEvent(ctx, eventId)
	if err != nil {
		return nil, err
	}

	return updateEvent, nil
}

// Delete implements EventUseCase.
func (eu *eventUseCase) Delete(c context.Context, eventId int) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(c, eu.timeout)
	defer cancel()

	deleteEvent, err := eu.repository.DeleteEvent(ctx, eventId)
	if err != nil {
		return nil, err
	}

	return deleteEvent, nil
}

func NewEventUseCase(eventRepo repository.EventRepository) EventUseCase {
	return &eventUseCase{
		repository: eventRepo,
	}
}
