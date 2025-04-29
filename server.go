package main

import (
	"log"
	"io"

	"github.com/emersion/go-smtp"
	"github.com/emersion/go-sasl"
	"gotempmail/sendmail"
)

type Backend struct{}

func (bkd *Backend) NewSession(state *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	auth bool
}

func (s *Session) AuthMechanisms() []string {
	return nil
}

func (s *Session) Auth(mech string) (sasl.Server, error) {
	return nil, nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
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

	go sendmail.NewMail()

	select {}
}