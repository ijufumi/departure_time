package service

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail"
	"go.uber.org/zap"
)

// SendMailService sends email
type SendMailService interface {
	Send(contents request.SendMail) error
}

type sendMailService struct {
	vendors []SendMailService
}

// Send allows sending email
func (service *sendMailService) Send(contents request.SendMail) error {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	var err error
	for _, vendor := range service.vendors {
		err = vendor.Send(contents)
		if err == nil {
			break
		}
		logger.Warn("Failed to send message", zap.Error(err))
	}
	return err
}

// NewSendMailService is factory method of SendMailService
func NewSendMailService(amazonSESService mail.AmazonSESService, sendGridService mail.SendGridService) SendMailService {
	return &sendMailService{vendors: []SendMailService{amazonSESService, sendGridService}}
}
