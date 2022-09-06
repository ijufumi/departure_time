package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/http/response"
	"github.com/ijufumi/email-service/internal/app/service"
	"net/http"
)

// SendMailHandler is the handler for sending mail
type SendMailHandler interface {
	Send(ctx *gin.Context)
}

type sendMailHandler struct {
	sendMailService service.SendMailService
}

// Send allows sending email via Email Vendors
func (handler *sendMailHandler) Send(ctx *gin.Context) {
	var contents request.SendMail
	if err := ctx.BindJSON(&contents); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Result{Status: false, Error: &err})
		return
	}
	if err := handler.sendMailService.Send(contents); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Result{Status: false, Error: &err})
		return
	}
	ctx.JSON(http.StatusOK, response.Result{Status: true})
}

// NewSendMailHandler is factory method of SendMailHandler
func NewSendMailHandler(sendMailService service.SendMailService) SendMailHandler {
	return &sendMailHandler{sendMailService: sendMailService}
}
