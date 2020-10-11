package sms

import "log"

type Service interface {
	Receive(from string, to string, message string)
}

type service struct {
}

func (s *service) Receive(from string, to string, message string) {
	log.Printf("New message from %s to %s: %s", from, to, message)
}

func NewService() Service {
	return &service{}
}
