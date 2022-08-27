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

func TestSendMailHandler_Send(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	jsonBody := strings.NewReader(`{"to_address":"to@test.com","from_address":"from@test.com","subject": "test subject", "body": "test body"}`)
	req, _ := http.NewRequest("POST", "/", jsonBody)
	ginContext.Request = req

	mockSendMailService := mock.NewMockSendMailService(false)
	sendMailHandler := NewSendMailHandler(mockSendMailService)
	sendMailHandler.Send(ginContext)

	asserts := assert.New(t)
	asserts.Equal(200, w.Result().StatusCode)
}
