package mail_vendor

type MailVendorInterface interface {
	Send(contents MailContents) error
}

type MailContents struct {
	ToAddress   string
	FromAddress string
	Contents    string
}
