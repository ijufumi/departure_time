package service

import "github.com/ijufumi/email-service/internal/app/service/vendors/mail"

type SendMailService interface {
	Send(contents mail.Contents) error
}

type sendMailService struct {
	vendors []SendMailService
}

func (service *sendMailService) Send(contents mail.Contents) error {
	var err error
	for _, vendor := range service.vendors {
		err = vendor.Send(contents)
		if err == nil {
			break
		}
	}
	return err
}

func NewSendMailService(vendors ...SendMailService) SendMailService {
	return &sendMailService{vendors: vendors}
}
