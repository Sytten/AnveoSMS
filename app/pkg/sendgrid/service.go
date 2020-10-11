package sendgrid

type Service interface {
	Send(content string, dst string)
}

type service struct {
}

func (s *service) Send(content string, dst string) {

}
