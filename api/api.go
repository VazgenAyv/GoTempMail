package api

import (
	"gotempmail/api/routes"

	"github.com/gin-gonic/gin"
)

func ApiServer() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8081")
}