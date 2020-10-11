package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/sytten/anveosms/pkg/server"
	"github.com/sytten/anveosms/pkg/sms"
)

func main() {
	logger, _ := zap.NewProduction()

	// Services
	ss := sms.NewService()

	// Server
	srv := server.New(ss, logger)

	// Start
	errs := make(chan error, 2)
	go func() {
		logger.Info("server started")
		errs <- http.ListenAndServe("0.0.0.0:9000", srv)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	<-errs
	logger.Info("server terminated")
}
