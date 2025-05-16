package api

import (
	"gotempmail/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func ApiServer() error {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 	  []string{"http://127.0.0.1:5173"},
		AllowMethods: 	  []string{"GET","POST", "PUT", "DELETE"},
		AllowHeaders: 	  []string{"Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(r)
	
	return r.Run(":8081")
}