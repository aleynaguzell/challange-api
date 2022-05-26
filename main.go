package main

import (
	"context"
	"github.com/aleynaguzell/getir-challange-api/controllers"
	"github.com/aleynaguzell/getir-challange-api/pkg/config"
	"github.com/aleynaguzell/getir-challange-api/pkg/logger"
	"github.com/aleynaguzell/getir-challange-api/pkg/mongo"
	"github.com/aleynaguzell/getir-challange-api/storage"
	"github.com/aleynaguzell/getir-challange-api/storage/memory"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func main() {
	logger.Init()
	config.Setup()

	mStorage := memory.New()
	mClient, err := mongo.Init()
	if err != nil {
		logger.Logger.Error("mongo connection error", err)
	}
	database := storage.New(mClient, mStorage)
	cf := controllers.NewControllerFactory(database)
	http.HandleFunc("/", HealthCheck)
	http.HandleFunc("/in-memory", cf.GetMemoryController().Get)
	http.HandleFunc("/in-memory/", cf.GetMemoryController().Set)
	http.HandleFunc("/records", cf.GetRecordController().GetRecords)

	httpServer := &http.Server{
		Addr: ":" + config.GetConfig().Server.Port,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {

			logger.Logger.Fatal("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	logger.Logger.Info("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		logger.Logger.Fatal("os.Kill - terminating...\n")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Logger.Error("error handled: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		logger.Logger.Info("stopped\n")
	}

	defer os.Exit(0)

}
