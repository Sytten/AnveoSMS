package email

import (
	"time"

	"go.uber.org/zap"
)

type loggingService struct {
	logger *zap.Logger
	next   Service
}

func NewLoggingService(logger *zap.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Send(content string) (err error) {
	defer func(begin time.Time) {
		s.logger.Info(
			"Email sent",
			zap.String("method", "Send"),
			zap.Duration("took", time.Since(begin)),
			zap.Error(err),
		)
	}(time.Now())
	return s.next.Send(content)
}
