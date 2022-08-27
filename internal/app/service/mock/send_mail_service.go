package mock

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service"
	"github.com/pkg/errors"
)

type mockSendMailService struct {
	hasError bool
}

func (mock *mockSendMailService) Send(contents request.SendMail) error {
	if mock.hasError {
		return errors.New("mock service error")
	}
	return nil
}

// NewMockSendMailService is factory method of Mock of SendMailService
func NewMockSendMailService(withError bool) service.SendMailService {
	return &mockSendMailService{hasError: withError}
}
