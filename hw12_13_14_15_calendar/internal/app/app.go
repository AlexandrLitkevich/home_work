package app

import (
	"context"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage"
)

type App struct { // TODO
}

type Logger interface { // TODO
}

type Storage interface {
	Add(ctx context.Context, title, desc string) (*storage.Event, error)
	Remove(ctx context.Context, id string) error
	Edit(ctx context.Context, event storage.Event) (*storage.Event, error)
	Get(ctx context.Context, id string) (*storage.Event, error)
	List(ctx context.Context) ([]storage.Event, error)
	Close() error
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
