package users

import (
	"strings"

	"github.com/thepranavjain/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Prepare() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
}

func (user *User) Validate() *errors.RestErr {
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

func (user *User) ValidatePassword() *errors.RestErr {
	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 8 {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
