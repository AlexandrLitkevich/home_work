package memorystorage

import (
	"context"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO sync test
func TestStorageAddEvent(t *testing.T) {
	store := New()

	ctx := context.TODO()

	event, err := store.Add(ctx, "this test event", "I' am test desc")
	require.NoError(t, err)
	require.NotNil(t, event)
}

func TestStorageRemoveEvent(t *testing.T) {
	store := New()

	ctx := context.TODO()

	event, err := store.Add(ctx, "this test event", "I' am test desc")
	require.NoError(t, err)
	require.NotNil(t, event)

	err = store.Remove(ctx, event.ID)
	require.NoError(t, err)
}

func TestStorageEditEvent(t *testing.T) {
	store := New()

	ctx := context.TODO()

	event, err := store.Add(ctx, "this test event", "I' am test desc")
	require.NoError(t, err)
	require.NotNil(t, event)

	const (
		editTitle = "this edit title"
		editDesc  = "this edit desc"
	)

	newEvent, err := store.Edit(ctx, storage.Event{ID: event.ID, Title: editTitle, Desc: editDesc})
	require.NoError(t, err)
	require.Equal(t, newEvent.Title, editTitle)
	require.Equal(t, newEvent.Desc, editDesc)
}

func TestStorageListEvents(t *testing.T) {
	store := New()

	ctx := context.TODO()

	for i := 0; i < 20; i++ {
		event, err := store.Add(ctx, "this test event", "I' am test desc")
		require.NoError(t, err)
		require.NotNil(t, event)
	}
	events := store.List(ctx)
	require.Equal(t, len(events), 20)
}

func TestStorageGetEvent(t *testing.T) {
	store := New()

	ctx := context.TODO()

	event, err := store.Add(ctx, "this test event", "I' am test desc")
	require.NoError(t, err)
	require.NotNil(t, event)

	getEvent, err := store.Get(ctx, event.ID)
	require.NoError(t, err)
	require.Equal(t, event, getEvent)

}
