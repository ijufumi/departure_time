package mail_vendor

type amazonSESService struct {
}

func (service *amazonSESService) Send(contents MailContents) error {
	return nil
}

func NewAmazonSESService() MailVendorInterface {
	return &amazonSESService{}
}
