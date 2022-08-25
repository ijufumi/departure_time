package mail

import (
	"github.com/ijufumi/email-service/internal/pkg/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridService struct {
	config *config.Config
}

func (service *sendGridService) Send(contents MailContents) error {
	toAddress := mail.NewEmail(contents.ToAddress, contents.ToAddress)
	fromAddress := mail.NewEmail(contents.FromAddress, contents.FromAddress)
	message := mail.NewSingleEmailPlainText(fromAddress, contents.Subject, toAddress, contents.Body)
	client := service.createClient()
	_, err := client.Send(message)
	return err
}

func (service *sendGridService) createClient() *sendgrid.Client {
	return sendgrid.NewSendClient(service.config.Mail.SendGrid.SendGridAPIKEY)
}

func NewSendGridService(config *config.Config) MailVendorInterface {
	return &sendGridService{
		config: config,
	}
}
