package mock

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail"
	"github.com/pkg/errors"
)

type mockAmazonSesService struct {
	hasError bool
}

// Send is mock of Send method
func (mock *mockAmazonSesService) Send(_ request.SendMail) error {
	if mock.hasError {
		return errors.New("mock mockAmazonSesService error")
	}

	return nil
}

// NewMockAmazonSesService is factory method of Mock of AmazonSESService
func NewMockAmazonSesService(hasError bool) mail.AmazonSESService {
	return &mockAmazonSesService{hasError}
}
