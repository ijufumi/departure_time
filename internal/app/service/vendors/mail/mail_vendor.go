package mail

import "github.com/ijufumi/email-service/internal/app/http/request"

// SendMailVendor is interface of mail vendor
type SendMailVendor interface {
	Send(contents request.SendMail) error
}
