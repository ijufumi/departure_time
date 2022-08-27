package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Health(t *testing.T) {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	healthHandler := NewHealthHandler()
	healthHandler.Health(ginContext)

	asserts := assert.New(t)
	asserts.Equal(200, w.Result().StatusCode)
}
