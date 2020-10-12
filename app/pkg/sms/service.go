package sms

import (
	"fmt"
	"github.com/sytten/anveosms/pkg/email"
	"go.uber.org/zap"
)

type Service interface {
	Receive(from string, to string, message string)
}

type service struct {
	email  email.Service
	logger *zap.Logger
}

func (s *service) Receive(from string, to string, message string) {
	// Create the content
	content := fmt.Sprintf("New message from %s to %s: %s", from, to, message)

	// Send the email
	err := s.email.Send(content)
	if err != nil {
		s.logger.Error("Unable to forward SMS")
	}
	// TODO: Implement retry with exponential backoff
}

func NewService(email email.Service, logger *zap.Logger) Service {
	return &service{
		email,
		logger,
	}
}
