package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user on the social media
type User struct {
	ID        uint64    `json: "id, omitempty"`
	Name      string    `json: "name, omitempty"`
	Nick      string    `json: "nick, omitempty"`
	Email     string    `json: "email, omitempty"`
	Password  string    `gorm:"size:100" json: "password, omitempty"`
	CreatedAt time.Time `json: "CreatedAt, omitempty" `
}

// Prepare will call validate and format methods on the user
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New(FieldisEmptyMessage("name"))
	}
	if user.Nick == "" {
		return errors.New(FieldisEmptyMessage("nick"))
	}
	if user.Email == "" {
		return errors.New(FieldisEmptyMessage("email"))
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}

	if step == CREATE && user.Password == "" {
		return errors.New(FieldisEmptyMessage("password"))
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == CREATE {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
