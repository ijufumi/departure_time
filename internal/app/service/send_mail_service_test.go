package service

import (
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMailService_Send(t *testing.T) {
	mockAmazonSesService := mock.NewMockAmazonSesService(false)
	mockSendGridService := mock.NewMockSendGridService(false)

	sendMailService := NewSendMailService(mockAmazonSesService, mockSendGridService)
	err := sendMailService.Send(request.SendMail{})

	asserts := assert.New(t)

	asserts.Nil(err)
}

func TestSendMailService_Send_Second_Success(t *testing.T) {
	mockAmazonSesService := mock.NewMockAmazonSesService(true)
	mockSendGridService := mock.NewMockSendGridService(false)

	sendMailService := NewSendMailService(mockAmazonSesService, mockSendGridService)
	err := sendMailService.Send(request.SendMail{})

	asserts := assert.New(t)

	asserts.Nil(err)
}

func TestSendMailService_First_Success(t *testing.T) {
	mockAmazonSesService := mock.NewMockAmazonSesService(false)
	mockSendGridService := mock.NewMockSendGridService(true)

	sendMailService := NewSendMailService(mockAmazonSesService, mockSendGridService)
	err := sendMailService.Send(request.SendMail{})

	asserts := assert.New(t)

	asserts.Nil(err)
}

func TestSendMailService_Error(t *testing.T) {
	mockAmazonSesService := mock.NewMockAmazonSesService(true)
	mockSendGridService := mock.NewMockSendGridService(true)

	sendMailService := NewSendMailService(mockAmazonSesService, mockSendGridService)
	err := sendMailService.Send(request.SendMail{})

	asserts := assert.New(t)

	asserts.NotNil(err)
	asserts.Equal(err.Error(), "mock mockSendGridService error")
}
