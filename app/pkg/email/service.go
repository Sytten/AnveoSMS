package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sytten/anveosms/pkg/config"
)

type Service interface {
	Send(content string) error
}

type service struct {
	client *sendgrid.Client
	config config.EmailConfiguration
}

func (s *service) Send(content string) error {
	// Prepare message
	from := mail.NewEmail("AnveoSMS", s.config.From)
	subject := "New SMS from AnveoSMS"
	to := mail.NewEmail("", s.config.To)
	message := mail.NewSingleEmail(from, subject, to, content, content)

	// Send message
	_, err := s.client.Send(message)
	return err
}

func NewService(config *config.Configuration) Service {
	return &service{
		client: sendgrid.NewSendClient(config.Email.SendgridApiKey),
		config: config.Email,
	}
}