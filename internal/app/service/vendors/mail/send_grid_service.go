package mail

import (
	"github.com/ijufumi/email-service/internal/pkg/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridService is mail service for SendGrid
type SendGridService interface {
	SendMailVendor
}

type sendGridService struct {
	config *config.Config
}

// Send allows sending mail via SendGrid
func (service *sendGridService) Send(contents Contents) error {
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

// NewSendGridService is factory method of SendGridService
func NewSendGridService(config *config.Config) SendGridService {
	return &sendGridService{
		config: config,
	}
}
