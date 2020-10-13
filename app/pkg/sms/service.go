package sms

import (
	"fmt"

	"github.com/sytten/anveosms/pkg/config"

	"github.com/sytten/anveosms/pkg/email"
	"go.uber.org/zap"
)

type Service interface {
	Receive(from string, to string, message string)
}

type service struct {
	email  email.Service
	config *config.Configuration
	logger *zap.Logger
}

func (s *service) Receive(from string, to string, message string) {
	// Create the content
	plainTextContent := fmt.Sprintf("New Incoming SMS\nFrom: %s\n%s", from, message)
	htmlContent, err := s.renderTemplate(from, message)
	if err != nil {
		htmlContent = plainTextContent
	}

	// Send the email
	err = s.email.Send(plainTextContent, htmlContent)
	if err != nil {
		s.logger.Error("Unable to forward SMS")
	}
	// TODO: Implement retry with exponential backoff
}

func NewService(email email.Service, config *config.Configuration, logger *zap.Logger) Service {
	return &service{
		email,
		config,
		logger,
	}
}
