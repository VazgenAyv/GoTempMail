package routes

import (
	"gotempmail/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	emailRoutes := r.Group("/emails")
	{
		emailRoutes.POST("/", handlers.ReceiveEmail)
		emailRoutes.GET("/all", handlers.GetMails)
		emailRoutes.GET("/email", handlers.GetMyMail)
	}
}