package httphandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"message_handler/internal/models"
	"net/http"
	"time"
)

// ...
func (h *HTTPMessageHandle) GetMessage(c *gin.Context) {
	message := &models.Message{}

	if err := c.BindJSON(&message); err != nil {
		FailOnErrorsHttp(c.Writer, err, "invalid input body", http.StatusBadRequest)
		return
	}

	err := h.htm.MessageService(context.Background(), message)
	if err != nil {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
	} else {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
	}
}

// ...
func (h *HTTPMessageHandle) Statistics(c *gin.Context) {
	dateString := c.Query("firstdate")
	firstDate, err := time.Parse("2006-01-02", dateString)
	FailOnErrorsHttp(c.Writer, err, "can't convert parse first time value", http.StatusBadRequest)

	dateString = c.Query("seconddate")
	secondDate, err := time.Parse("2006-01-02", dateString)
	FailOnErrorsHttp(c.Writer, err, "can't convert parse second time value", http.StatusBadRequest)

	stat := &models.Statistics{FirstDate: firstDate, SecondDate: secondDate}

	a, err := h.htm.StatisticsService(context.Background(), stat)
	if err != nil {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
	} else {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Connection:", "close")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, a)
	}
}
