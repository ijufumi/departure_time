package mail

// SendMailVendor is interface of mail vendor
type SendMailVendor interface {
	Send(contents Contents) error
}
