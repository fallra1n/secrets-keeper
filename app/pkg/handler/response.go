package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
}

type statusResponse struct {
	Status string `json:"status"`
}
