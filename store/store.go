package store

import (
	"gotempmail/models"
	"sync"
)

var Store = &MailStore{}

type MailStore struct {
	sync.RWMutex
	Emails []models.Email
}

func(m *MailStore) Add(email models.Email) {
	m.Lock()
	defer m.Unlock()
	m.Emails = append(m.Emails, email)
}

func(m *MailStore) GetAll() []models.Email {
	m.RLock()
	defer m.RUnlock()
	return append([]models.Email(nil), m.Emails...)
}
