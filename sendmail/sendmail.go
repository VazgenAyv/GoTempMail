package sendmail

import (
	"log"
	"net/smtp"
)

func NewMail() {
	server := "localhost:1025"
	from := "sender@example.com"
	to := "receiver@example.com"
	subject := "Test Subject"
	body := "This is a test email"

	message := []byte("Subject: " + subject + "\r\n" + 
		"From: " + from + "\r\n" + 
		"To: " + to + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(
		server,
		nil,
		from,
		[]string{to},
		message,
	)
	
	if err != nil {
		log.Fatal("Failed to send email: ", err)
	} else {
		log.Println("Email sent successfully")
	}
}