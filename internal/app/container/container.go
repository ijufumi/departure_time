package container

import (
	"github.com/ijufumi/email-service/internal/app/http/handler"
	"github.com/ijufumi/email-service/internal/app/service"
	"github.com/ijufumi/email-service/internal/app/service/vendors/mail"
	"github.com/ijufumi/email-service/internal/pkg/config"
	"go.uber.org/dig"
)

// Container is DI container
type Container interface {
	GetHealthHandler() handler.HealthHandler
	GetSendMailHandler() handler.SendMailHandler
}

type container struct {
	container *dig.Container
}

func (c *container) provide() {
	// config
	_ = c.container.Provide(config.Load)

	// service
	_ = c.container.Provide(service.NewSendMailService)
	_ = c.container.Provide(mail.NewAmazonSESService)
	_ = c.container.Provide(mail.NewSendGridService)

	// handler
	_ = c.container.Provide(handler.NewHealthHandler)
	_ = c.container.Provide(handler.NewSendMailHandler)
}

// GetHealthHandler returns instance of HealthHandler
func (c *container) GetHealthHandler() handler.HealthHandler {
	var healthHandler handler.HealthHandler
	_ = c.container.Invoke(func(_healthHandler handler.HealthHandler) {
		healthHandler = _healthHandler
	})
	return healthHandler
}

// GetSendMailHandler returns instance of SendMailHandler
func (c *container) GetSendMailHandler() handler.SendMailHandler {
	var sendMailHandler handler.SendMailHandler
	_ = c.container.Invoke(func(_sendMailHandler handler.SendMailHandler) {
		sendMailHandler = _sendMailHandler
	})
	return sendMailHandler
}

// NewContainer creates instance of Container
func NewContainer() Container {
	c := container{container: dig.New()}
	c.provide()
	return &c
}
