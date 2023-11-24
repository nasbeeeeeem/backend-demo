package repository

import (
	"backend-demo/ent"
	"context"
)

type EventRepository interface {
	CreateEvent(c context.Context, name string, userId int) (*ent.Event, error) //POST:/event
	GetEvents(c context.Context) ([]*ent.Event, error)                          //GET:/event
	GetEventsById(c context.Context, userId int) ([]*ent.Event, error)          //GET:/event/:id
	UpdateEvent(c context.Context, eventId int) (*ent.Event, error)             //PUT:/event/:id
	DeleteEvent(c context.Context, eventId int) (*ent.Event, error)             //DELETE:/event/:id
}
