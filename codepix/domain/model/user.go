package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
	"regexp"
) 

const (
	LeftLength    number = 3
	RightLength   number = 50
	
)

type User struct{ 
	Base      `valid:"required"`
	Name      string     `json:"name" valid:"notnull"`
	Email     string     `json:"code" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if isEmailValid(user.Email) {
	    return errors.New("invalid email for the user")
	}
	
	if err != nil {
		return err
	}
	return nil
}

func NewUser( name string,email string) (*User, error) {
	user := User{
		Name: name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func isEmailValid(e string) bool {
	if len(e) < LeftLength && len(e) > RightLength{
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(e)
}