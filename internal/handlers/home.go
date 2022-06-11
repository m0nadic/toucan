package handlers

import (
	"github.com/gin-gonic/gin"
	"toucan/internal/assets"
)

type HomeHandler struct {
}

func NewHomeHandler() HomeHandler {
	return HomeHandler{}
}
func (hh HomeHandler) Index(ctx *gin.Context) {
	data, _ := assets.Asset("assets/index.html")
	ctx.Writer.Header().Set("Content-Type", "text/html")
	_, _ = ctx.Writer.Write(data)
}
