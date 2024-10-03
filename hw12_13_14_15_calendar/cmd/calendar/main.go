package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/cmd"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/jackc/pgx/v5"
)

func main() {
	cmd.Execute()

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("fail to read config")
		os.Exit(1)
	}

	ctx := context.Background()

	appLogger := logger.New()
	var storage app.Storage
	appLogger.Info("the logger has been successfully configured", "config", cfg)
	if cfg.Storage.StorageType == "memory" {
		storage = memorystorage.New()
	} else if cfg.Storage.StorageType == "sql" {
		appLogger.Info("connection to database....")
		appLogger.Info("cfg.Storage.Postrges.Url", "url", cfg.Storage.Postgres.Url)

		conn, err := pgx.Connect(ctx, cfg.Storage.Postgres.Url) // TODO check url
		if err != nil {
			slog.Error("fail to connection postgres")
		}

		appLogger.Info("connection to database success")

		// err = conn.Ping(ctx)
		// if err != nil {
		// 	appLogger.Error("failed connection database", "error", err)
		// }

		storage = sqlstorage.New(appLogger, cfg, conn)
		defer storage.Close(ctx)

		err = storage.Ping(ctx)
		if err != nil {
			appLogger.Error("failed connection database", "error", err)
		}

		//var greeting string
		//err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		//	os.Exit(1)
		//}
		//appLogger.Warn("this test greeting", greeting)
		//slog.Warn("this test greeting", "GRETING", greeting)
	}

	appLogger.Info("create database")
	calendar := app.New(appLogger, storage)

	server := internalhttp.NewServer(appLogger, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			appLogger.Error("failed to stop http server: " + err.Error())
		}
	}()

	appLogger.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		appLogger.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
