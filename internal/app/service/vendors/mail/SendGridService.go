package mail

import "github.com/ijufumi/email-service/internal/pkg/config"

type sendGridService struct {
	config *config.Config
}

func (service *sendGridService) Send(contents MailContents) error {
	return nil
}

func NewSendGridService(config *config.Config) MailVendorInterface {
	return &sendGridService{
		config: config,
	}
}
