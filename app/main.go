package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sytten/anveosms/pkg/config"
	"github.com/sytten/anveosms/pkg/email"

	"go.uber.org/zap"

	"github.com/sytten/anveosms/pkg/server"
	"github.com/sytten/anveosms/pkg/sms"
)

var BuildVersion = "development"

func main() {
	logger, _ := zap.NewProduction()
	logger = logger.Named("anveoSms")

	// Config
	configuration, err := config.NewConfiguration(logger.Named("config"))
	if err != nil {
		logger.Error("Unable to load the configuration", zap.Error(err))
		return
	}

	// Services
	emailService := email.NewLoggingService(
		logger.Named("email"),
		email.NewService(configuration),
	)
	smsService := sms.NewService(emailService, configuration, logger.Named("sms"))

	// Server
	srv := server.New(smsService, logger)

	// Start
	errs := make(chan error, 2)
	go func() {
		logger.Info("server started", zap.String("version", BuildVersion))
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
