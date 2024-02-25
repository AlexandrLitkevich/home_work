package memorystorage

import (
	"context"
	"errors"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"sync"
)

type Memory struct {
	events map[uuid.UUID]storage.Event
	mu     sync.RWMutex //nolint:unused
}

func New() *Memory {
	return &Memory{
		events: make(map[uuid.UUID]storage.Event),
	}
}

func (m *Memory) Add(_ context.Context, title string, desc string) (*storage.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	if _, ok := m.events[key]; ok {
		return nil, errors.New("invalid uuid")
	}

	event := storage.Event{
		ID:    key.String(),
		Title: title,
		Desc:  desc,
	}

	m.events[key] = event

	newEvent := m.events[key]

	return &newEvent, nil
}
func (m *Memory) Remove(_ context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := uuid.MustParse(id)

	delete(m.events, key)

	return nil
}
func (m *Memory) Edit(_ context.Context, event storage.Event) (*storage.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := uuid.MustParse(event.ID)

	m.events[key] = event

	return &event, nil
}
func (m *Memory) Get(_ context.Context, id string) (*storage.Event, error) {
	key := uuid.MustParse(id)

	event, ok := m.events[key]
	if !ok {
		return nil, errors.New("event not found")
	}
	return &event, nil
}
func (m *Memory) List(_ context.Context) ([]storage.Event, error) {
	response := make([]storage.Event, 0, len(m.events))

	for _, event := range m.events {
		response = append(response, event)
	}
	return response, nil
}

func (m *Memory) Close() error {
	return nil
}
