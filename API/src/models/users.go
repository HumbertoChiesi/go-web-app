package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//User represents a user using the social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
}

//Prepare will call the methods to validate and format the received user
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("the name field cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("the nick field cannot be empty")
	}

	if user.Email == "" {
		return errors.New("the email field cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("this email is not valid")
	}

	if stage == "registering" && user.Password == "" {
		return errors.New("the password field cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if stage == "registering" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashPassword)
	}
	return nil
}
