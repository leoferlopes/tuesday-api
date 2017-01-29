package api

import (
	"github.com/leoferlopes/tuesday-api/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/leoferlopes/tuesday-api/configuration"
)

// Run define the handled endpoint on API.
func Run() {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/menu", routes.GetMenu(configuration.GetMenu()))

	r.Run(configuration.GetAddress())
}
