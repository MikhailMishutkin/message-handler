package httphandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"message_handler/internal/models"
)

type HTTPMessageHandle struct {
	htm HTTPMessageManager
}

func NewHTTPMessageHandle(htm HTTPMessageManager) *HTTPMessageHandle {
	return &HTTPMessageHandle{htm: htm}
}

type HTTPMessageManager interface {
	MessageService(context.Context, *models.Message) error
	StatisticsService(context.Context, *models.Statistics) (*models.Statistics, error)
}

func (h *HTTPMessageHandle) RegisterMessage(router *gin.Engine) {
	router.POST("/message", h.GetMessage)
	router.GET("/statistics", h.Statistics)
}
