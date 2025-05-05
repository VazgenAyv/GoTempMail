package handlers

import (
	"gotempmail/models"
	"gotempmail/store"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var request models.Email

func ReceiveEmail(c *gin.Context) {

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return 
	}

	log.Printf("Receiving email from: %s to: %s, Subject: %s\n",request.From, request.To, request.Subject)

	c.JSON(http.StatusOK, gin.H{"message": "Email received"})
}

func GetMails(c *gin.Context) {
	emails := store.Store.GetAll()
	c.JSON(http.StatusOK, emails)
}

func GetMyMail (c *gin.Context) {
	email := c.Query("email")

	var filtered []models.Email

	var emails = store.Store.GetAll()
	for _, mail := range  emails {
		if email == mail.To {
			filtered = append(filtered, mail)
		}
	}
	c.JSON(http.StatusOK, filtered)
}