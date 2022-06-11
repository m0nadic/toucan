package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toucan/internal/handlers"
	"toucan/internal/service"
)

func StartHTTP(host string, port int) error {
	r := gin.Default()

	r.GET("/", handlers.NewHomeHandler().Index)
	r.POST("/build", handlers.NewBuildHandler(service.NewBuilderService()).Build)

	return r.Run(fmt.Sprintf("%s:%d", host, port))
}
