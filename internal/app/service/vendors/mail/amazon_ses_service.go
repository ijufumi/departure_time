package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/ijufumi/email-service/internal/pkg/config"
)

type amazonSESService struct {
	config *config.Config
}

func (service *amazonSESService) Send(contents MailContents) error {
	svc, err := service.createSesService()
	if err != nil {
		return err
	}

	input := ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(contents.ToAddress),
			},
			CcAddresses: []*string{},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data:    aws.String(contents.Body),
					Charset: aws.String(service.config.Mail.Charset),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(contents.Subject),
				Charset: aws.String(service.config.Mail.Charset),
			},
		},
	}

	_, err = svc.SendEmail(&input)
	return err
}

func (service *amazonSESService) createSesService() (*ses.SES, error) {
	sesSession, err := session.NewSession(&aws.Config{
		Region: aws.String(service.config.Mail.SES.AwsRegion),
	})
	if err != nil {
		return nil, err
	}

	return ses.New(sesSession), nil
}

func NewAmazonSESService(config *config.Config) MailVendorInterface {
	return &amazonSESService{config: config}
}
