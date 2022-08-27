package mock

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail"
	"github.com/pkg/errors"
)

type mockSendGridService struct {
	hasError bool
}

// Send is mock of Send method
func (mock *mockSendGridService) Send(_ request.SendMail) error {
	if mock.hasError {
		return errors.New("mock mockSendGridService error")
	}

	return nil
}

// NewMockSendGridService is factory method of Mock of SendGridService
func NewMockSendGridService(hasError bool) mail.SendGridService {
	return &mockSendGridService{hasError}
}
