package main

import (
	"context"
	"flag"
	internalapp "github.com/socialdistance/spa-test/internal/app"
	internalconfig "github.com/socialdistance/spa-test/internal/config"
	internallogger "github.com/socialdistance/spa-test/internal/logger"
	internalhttp "github.com/socialdistance/spa-test/internal/server/http"
	sqlstorage "github.com/socialdistance/spa-test/internal/storage/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := internalconfig.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed load config %s", err)
	}

	logg, err := internallogger.New(config.Logger)
	if err != nil {
		log.Fatalf("Failed load logger %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	store := CreateStorage(ctx, *config)
	spa := internalapp.New(logg, store)

	server := internalhttp.NewServer(logg, spa, config.HTTP.Host, config.HTTP.Port)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("spa is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func CreateStorage(ctx context.Context, config internalconfig.Config) internalapp.Storage {
	var store internalapp.Storage

	switch config.Storage.Type {
	case internalconfig.SQL:
		sqlStore := sqlstorage.New(ctx, config.Storage.URL)
		err := sqlStore.Connect(ctx)
		if err != nil {
			log.Fatalf("Unable to connect database: %s", err)
		}
		store = sqlStore
	default:
		log.Fatalf("Dont know type storage: %s", config.Storage.Type)
	}

	return store
}
