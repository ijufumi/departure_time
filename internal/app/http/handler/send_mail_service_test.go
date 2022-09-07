package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/service/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestSendMailHandler_Send is normal test case
func TestSendMailHandler_Send(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	jsonBody := strings.NewReader(`{"to_address":"to@test.com","subject": "test subject", "body": "test body"}`)
	req, _ := http.NewRequest("POST", "/", jsonBody)
	ginContext.Request = req

	mockSendMailService := mock.NewMockSendMailService(false)
	sendMailHandler := NewSendMailHandler(mockSendMailService)
	sendMailHandler.Send(ginContext)

	asserts := assert.New(t)
	asserts.Equal(200, w.Result().StatusCode)
}

// TestSendMailHandler_Send_ValidationError is simple validation error
func TestSendMailHandler_Send_ValidationError(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	jsonBody := strings.NewReader(`{}`)
	req, _ := http.NewRequest("POST", "/", jsonBody)
	ginContext.Request = req

	mockSendMailService := mock.NewMockSendMailService(false)
	sendMailHandler := NewSendMailHandler(mockSendMailService)
	sendMailHandler.Send(ginContext)

	asserts := assert.New(t)
	asserts.Equal(400, w.Result().StatusCode)
}

// TestSendMailHandler_Send_ValidationMailFormatError is email format error
func TestSendMailHandler_Send_ValidationMailFormatError_To(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	jsonBody := strings.NewReader(`{"to_address":"to@test","subject": "test subject", "body": "test body"}`)
	req, _ := http.NewRequest("POST", "/", jsonBody)
	ginContext.Request = req

	mockSendMailService := mock.NewMockSendMailService(false)
	sendMailHandler := NewSendMailHandler(mockSendMailService)
	sendMailHandler.Send(ginContext)

	asserts := assert.New(t)
	asserts.Equal(400, w.Result().StatusCode)
}

// TestSendMailHandler_Send_Error is error caused by send mail service
func TestSendMailHandler_Send_Error(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	jsonBody := strings.NewReader(`{"to_address":"to@test.jp","subject": "test subject", "body": "test body"}`)
	req, _ := http.NewRequest("POST", "/", jsonBody)
	ginContext.Request = req

	mockSendMailService := mock.NewMockSendMailService(true)
	sendMailHandler := NewSendMailHandler(mockSendMailService)
	sendMailHandler.Send(ginContext)

	asserts := assert.New(t)
	asserts.Equal(500, w.Result().StatusCode)
}
