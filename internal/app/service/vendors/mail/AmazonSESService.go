package mail

type amazonSESService struct {
}

func (service *amazonSESService) Send(contents MailContents) error {
	return nil
}

func NewAmazonSESService() MailVendorInterface {
	return &amazonSESService{}
}
