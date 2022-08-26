package mail

type SendMailVendor interface {
	Send(contents Contents) error
}

type Contents struct {
	ToAddress   string
	FromAddress string
	Subject     string
	Body        string
}
