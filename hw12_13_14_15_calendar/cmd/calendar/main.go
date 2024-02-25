package main

import (
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/cmd"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/logger"
	"log/slog"
)

func main() {
	cmd.Execute()

	_, err := config.NewConfig()
	if err != nil {
		slog.Error("fail to read config")
		panic(err)
	}

	appLogger := logger.New()

	appLogger.Info("the logger has been successfully configured")
	//if cfg.Storage.ty

	//
	//storage := memorystorage.New()
	//calendar := app.New(appLogger, storage)
	//
	//server := internalhttp.NewServer(appLogger, calendar)
	//
	//ctx, cancel := signal.NotifyContext(context.Background(),
	//	syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	//defer cancel()
	//
	//go func() {
	//	<-ctx.Done()
	//
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//	defer cancel()
	//
	//	if err := server.Stop(ctx); err != nil {
	//		appLogger.Error("failed to stop http server: " + err.Error())
	//	}
	//}()
	//
	//appLogger.Info("calendar is running...")
	//
	//if err := server.Start(ctx); err != nil {
	//	appLogger.Error("failed to start http server: " + err.Error())
	//	cancel()
	//	os.Exit(1) //nolint:gocritic
	//}
}
