package service

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail"
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
	var err error
	for _, vendor := range service.vendors {
		err = vendor.Send(contents)
		if err == nil {
			break
		}
	}
	return err
}

// NewSendMailService is factory method of SendMailService
func NewSendMailService(amazonSESService mail.AmazonSESService, sendGridService mail.SendGridService) SendMailService {
	return &sendMailService{vendors: []SendMailService{amazonSESService, sendGridService}}
}
