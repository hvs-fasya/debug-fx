package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/logger"
)

var v1handlers = []handler{
	{Method: http.MethodGet, Path: "/lessons", HandlerFunc: GetLessons},
}

type handler struct {
	Method      string
	Path        string
	HandlerFunc func(*gin.Context, logger.Logger)
}

func GetLessons(ctx *gin.Context, logger logger.Logger) {
	logger.Info("GetLessons")
	return
}
