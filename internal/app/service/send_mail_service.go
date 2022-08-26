package service

import "github.com/ijufumi/email-service/internal/app/service/vendors/mail"

type SendMailServiceInterface interface {
	Send(contents mail.MailContents) error
}

type sendMailService struct {
	vendors []SendMailServiceInterface
}

func (service *sendMailService) Send(contents mail.MailContents) error {
	var err error
	for _, vendor := range service.vendors {
		err = vendor.Send(contents)
		if err == nil {
			break
		}
	}
	return err
}

func NewSendMailService(vendors ...SendMailServiceInterface) SendMailServiceInterface {
	return &sendMailService{vendors: vendors}
}
