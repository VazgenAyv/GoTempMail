package handlers

import (
	"fmt"
	"gotempmail/models"
	"gotempmail/store"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func randomString(length int) string {
	b := strings.Builder{}
	for i := 0; i < length; i++ {
		b.WriteByte(charset[rand.Intn(len(charset))])
	}
	return b.String()
}

func randomEmail() string {
	user := randomString(6)
	domain := randomString(6)
	tld := []string{"com", "net", "org", "io"}[rand.Intn(4)]
	return fmt.Sprintf("%s@%s.%s", user, domain, tld)
}

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

func GetEmail (c *gin.Context) {
	newEmail := randomEmail()
	log.Println(newEmail)
	c.JSON(http.StatusOK, gin.H{"email": newEmail})
}
