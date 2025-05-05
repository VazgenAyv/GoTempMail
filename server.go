package main

import (
	"bytes"
	"encoding/json"
	"gotempmail/api"
	"gotempmail/models"
	"gotempmail/store"
	"io"
	"log"
	"net/http"
	"net/mail"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
)


type Backend struct{}


func (bkd *Backend) NewSession(state *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	from string
	to 	 string
}

func (s *Session) AuthMechanisms() []string {
	return nil
}

func (s *Session) Auth(mech string) (sasl.Server, error) {
	return nil, nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	s.to = to
	return nil
}

func ( s *Session) Data(r io.Reader) error {

	raw, err := io.ReadAll(r)
	if err != nil {
		log.Println("Failed to read message:", err)
		return err
	}

	msg, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		log.Println("Failed to parse headers: ", err)
		return err
	}

	subject := msg.Header.Get("Subject")
	date := msg.Header.Get("Date")

	env, err := enmime.ReadEnvelope(bytes.NewReader(raw))
	if err != nil {
		log.Println("Failed to parse MIME envelope: ", err)
		return err
	}

	mail := models.Email{
		From: 	 s.from,
		To:		 s.to,
		Subject: subject,
		Body: 	 env.Text,
		Date: 	 date,
	}
	jsonData, err := json.Marshal(mail)

	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8081/emails/", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println("Error sending to API: ", err)
	}

	store.Store.Add(mail)

	defer resp.Body.Close()

	log.Println("Data send successfully", resp.StatusCode)

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func startServer() {
	be := &Backend{}
	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "127.0.0.1"
	s.AllowInsecureAuth = true

	log.Println("SMTP Running at", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	go startServer()

	go api.ApiServer()

	select {}
}