package main

import (
	httpcontroller "darkness8129/validation-service/app/controller/http"
	"darkness8129/validation-service/config"
	"darkness8129/validation-service/packages/httpserver"
	"darkness8129/validation-service/packages/logging"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to get config", "err", err)
	}

	// init http server and start it
	httpServer := httpserver.NewGinHTTPServer(httpserver.Options{
		Addr:         cfg.HTTP.Addr,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
	})

	router, ok := httpServer.Router().(*gin.Engine)
	if !ok {
		logger.Fatal("failed type assertion for router")
	}

	httpcontroller.New(httpcontroller.Options{
		Router: router,
		Logger: logger,
	})

	httpServer.Start()

	// graceful shutdown the http server with a timeout
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app interrupt", "signal", s.String())
	case err := <-httpServer.Notify():
		logger.Error("err from notify ch", "err", err)
	}

	err = httpServer.Shutdown(cfg.ShutdownTimeout)
	if err != nil {
		logger.Error("failed to shutdown server", "err", err)
	}
}
