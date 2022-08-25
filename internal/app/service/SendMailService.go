package service

import "github.com/ijufumi/email-service/internal/app/service/vendors/mail"

type SendMailServiceInterface interface {
	Send(contents mail.MailContents) error
}

type sendMailService struct {
}

func (service *sendMailService) Send(contents mail.MailContents) error {
	return nil
}

func NewSendMailService() SendMailServiceInterface {
	return &sendMailService{}
}
