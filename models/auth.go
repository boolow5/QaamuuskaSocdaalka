package models

import (
	//"gopkg.in/mgo.v2"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/boolow5/bolow/encrypt"
)

// authentcation models
// 1. user, 2. profile
// OVERVIEW: users sign up and get access to the reader/writer privileges.
// Admin users must be put in the database by another admin

type User struct {
	Id        int       `json:"user_id" orm:"auto"`
	Username  string    `json:"username" orm:"unique;size(30)"`
	Password  string    `json:"password" orm:"size(100)"`
	Role      string    `json:"role" orm:"size(20)"`
	Profile   *Profile  `json:"profile" orm:"rel(one);on_delete(cascade)"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *User) TableName() string {
	return "user"
}

func (this *User) Valid() bool {
	return (len(this.Username) > 2 && len(this.Password) > 4 && this.Profile.FirstName != "")
}

func (this User) String() string {
	return this.Username
}

type Profile struct {
	Id              int    `json:"profile_id" orm:"auto"`
	User            *User  `json:"user" orm:"reverse(one)"`
	FirstName       string `json:"first_name" orm:"size(30)"`
	MiddleName      string `json:"middle_name" orm:"size(30)"`
	LastName        string `json:"last_name" orm:"size(30)"`
	AnswerCount     int    `json:"answer_count" `
	QuestionCount   int    `json:"question_count" `
	AnswerViewCount int    `json:"anwer_view_count" `
}

func (this *Profile) TableName() string {
	return "profile"
}

// Authenticate checks the username and password and returns error if the authentication fails.
func (this *User) Authenticate() (error, *User) {
	if len(this.Username) < 2 && len(this.Password) < 2 {
		return errors.New("NotEnoughDataError: Username and password are required"), this
	}

	err, user := GetUserByUsername(this.Username)
	if err != nil {
		return err, this
	}
	if user.Password == "" {
		return errors.New("Invalid username or password"), this
	}

	correct, err := encrypt.AuthenticatePassword(user.Password, this.Password)

	if err != nil {
		return err, this
	}

	if !correct {
		return errors.New("Incorrect username or password"), this
	}
	user.Password = "[hidden-for-security-reasons]"
	return nil, user
}

func (this *User) Authorize(role string) (bool, error) {
	err, user := GetUserByUsername(this.Username)
	if err != nil {
		return false, err
	}

	fmt.Println("\n\nUser Role: ", user.Role)

	if role == "*" {
		return true, nil
	}

	if strings.ToLower(user.Role) != strings.ToLower(role) {
		return false, errors.New("You're not authorized, because your role is not " + role)
	}
	return true, nil
}
