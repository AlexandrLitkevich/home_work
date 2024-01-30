package main

import (
	"fmt"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/cmd"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/logger"
	"log/slog"
)

func main() {
	cmd.Execute()

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("fail to read config")
		panic(err)
	}

	fmt.Println("This my cfg", slog.String("This CFG", cfg.Logger.Level))

	logg := logger.New(cfg.Logger.Level)

	logg.Warn("this my implemints logg")
	//
	//storage := memorystorage.New()
	//calendar := app.New(logg, storage)
	//
	//server := internalhttp.NewServer(logg, calendar)
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
	//		logg.Error("failed to stop http server: " + err.Error())
	//	}
	//}()
	//
	//logg.Info("calendar is running...")
	//
	//if err := server.Start(ctx); err != nil {
	//	logg.Error("failed to start http server: " + err.Error())
	//	cancel()
	//	os.Exit(1) //nolint:gocritic
	//}
}
