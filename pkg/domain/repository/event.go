package repository

import (
	"backend-demo/pkg/domain/model"
	"context"
)

type EventRepository interface {
	CreateEvent(c context.Context, name string, userId int) (*model.Event, error) //POST:/event
	GetEvents(c context.Context) ([]*model.Event, error)                          //GET:/event
	GetEventsById(c context.Context, userId int) ([]*model.Event, error)          //GET:/event/:id
	UpdateEvent(c context.Context, eventId int) (*model.Event, error)             //PUT:/event/:id
	DeleteEvent(c context.Context, eventId int) (*model.Event, error)             //DELETE:/event/:id
}
