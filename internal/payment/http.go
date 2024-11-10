package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) RegisterRoutes(c *gin.Engine) {
	c.POST("/api/webhook", h.HandleWebhook)
}

func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	logrus.Info("got webhook from stripe")
}
