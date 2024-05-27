package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/carlos1377/devbook/api/authentication"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}
	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name field should not be empty")
	}
	if u.Nick == "" {
		return errors.New("nick field should not be empty")
	}
	if u.Email == "" {
		return errors.New("email field should not be empty")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email field not valid")
	}

	if step == "create" && u.Password == "" {
		return errors.New("password field should not be empty")
	}
	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "create" {
		hashedPassword, err := authentication.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}
