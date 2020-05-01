package services

import (
	"github.com/thepranavjain/bookstore_users-api/domain/users"
	"github.com/thepranavjain/bookstore_users-api/utils/date_utils"
	"github.com/thepranavjain/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user Id")
	}
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.Prepare()
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	user.Prepare()
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
