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
	return nil
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
