package sqlstorage

import (
	"context"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v5"
)

//https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

type Storage struct {
	db  *pgx.Conn
	cfg config.Config
	log Logger
}

type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func New(log Logger, cfg config.Config, connect *pgx.Conn) *Storage {
	return &Storage{
		cfg: cfg,
		log: log,
		db:  connect,
	}
}

func (s *Storage) Ping(ctx context.Context) error {
	return s.db.Ping(ctx)
}

//func (s *Storage) Connect(ctx context.Context) error {
//	conn, err := pgx.Connect(ctx, s.cfg.Storage.Postgres.Url)
//	if err != nil {
//		return err
//	}
//
//	s.db = conn
//	return nil
//}

func (s *Storage) Add(ctx context.Context, title, desc string) (*storage.Event, error) {
	//query := `INSERT INTO events (id,title, desc ) VALUES (@id, @title, @desc)`
	//args := pgx.NamedArgs{
	//	"id":    "123",
	//	"title": "this tilte",
	//	"desc":  "THis my decription",
	//}
	//
	//_, err := s.db.Exec(ctx, query, args)
	//if err != nil {
	//	return err
	//}

	return nil, nil
}

func (s *Storage) Get(ctx context.Context, id string) (*storage.Event, error) {
	// TODO
	return nil, nil
}

func (s *Storage) List(ctx context.Context) ([]storage.Event, error) {
	list := make([]storage.Event, 0)
	return list, nil
}

func (s *Storage) Edit(ctx context.Context, event storage.Event) (*storage.Event, error) {
	// TODO
	return nil, nil
}

func (s *Storage) Remove(ctx context.Context, id string) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.db.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}
