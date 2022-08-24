package mail_vendor

type sendGridService struct {
}

func (service *sendGridService) Send(contents MailContents) error {
	return nil
}

func NewSendGridService() MailVendorInterface {
	return &sendGridService{}
}
