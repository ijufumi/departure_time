package mail

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/pkg/config"
)

// AmazonSESService is mail service for Amazon SES
type AmazonSESService interface {
	SendMailVendor
}

type amazonSESService struct {
	config *config.Config
}

// Send allows sending mail via SES
func (service *amazonSESService) Send(contents request.SendMail) error {
	svc, err := service.createSesService()
	if err != nil {
		return err
	}

	input := sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				contents.ToAddress,
			},
			CcAddresses: []string{},
		},
		FromEmailAddress: aws.String(service.config.Mail.FromAddress),
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data:    aws.String(contents.Body),
						Charset: aws.String(service.config.Mail.Charset),
					}},
				Subject: &types.Content{
					Data:    aws.String(contents.Subject),
					Charset: aws.String(service.config.Mail.Charset),
				},
			},
		},
	}

	_, err = svc.SendEmail(context.Background(), &input)
	return err
}

func (service *amazonSESService) createSesService() (*sesv2.Client, error) {
	awsConfiguration, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	return sesv2.NewFromConfig(awsConfiguration), nil
}

// NewAmazonSESService is factory method of AmazonSESService
func NewAmazonSESService(config *config.Config) AmazonSESService {
	return &amazonSESService{config: config}
}
