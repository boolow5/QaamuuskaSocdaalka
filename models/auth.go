package models

import (
	//"gopkg.in/mgo.v2"
	"errors"
	"fmt"
	"log"
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
	if len(this.Username) < 2 {
		fmt.Println("Short username")
		return false
	} else if len(this.Password) < 4 {
		fmt.Println("Short password")
		return false
	} else if len(this.Profile.FirstName) < 1 {
		fmt.Println("Short first name")
		return false
	}
	return true
}

func (this User) String() string {
	return this.Username
}

type Profile struct {
	Id         int    `json:"profile_id" orm:"auto"`
	User       *User  `json:"user" orm:"reverse(one)"`
	FirstName  string `json:"first_name" orm:"size(30)"`
	MiddleName string `json:"middle_name" orm:"size(30)"`
	LastName   string `json:"last_name" orm:"size(30)"`
}

func (this *Profile) TableName() string {
	return "profile"
}

// Add adds new user to the database, and returns error or false if adding is failed
func (this *User) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("ValidationError: fill all the required fields")
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = time.Now()

	if saved, err := CreateUser(this); err != nil || !saved {
		return saved, err
	}

	return true, nil
}

// Update chenges the modified fields of the object and ignores the emtpy ones.
func (this *User) Update() (bool, error) {

	if this.Id < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &User{Id: this.Id}
	oldItem.Profile = this.Profile
	updated, err := UpdateItem(oldItem, this)
	if err != nil {
		return false, err
	}
	if !updated {
		return false, nil
	}

	return true, nil
}

// Delete removes a user from the database
func (this *User) Delete() (bool, error) {
	if this.Id < 1 {
		return false, errors.New("IntegrityError: invalid user_id")
	}

	if deleted, err := DeleteItem(this); err != nil || !deleted {
		return deleted, err
	}

	return true, nil
}

// SetPassword is the recommended way when changing password because passwords need to be hashed, they cannot be plain text.
func (this *User) SetPassword(newPassword string) error {
	var err error
	this.Password, err = encrypt.HashPassword(newPassword)
	if err != nil {
		return err
	}
	if !encrypt.IsHash(this.Password) {
		return errors.New("Failed to hash this password")
	}
	return nil
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
		log.Println("Error from encrypt")
		return err, this
	}

	if !correct {
		return errors.New("Incorrect username or password"), this
	}

	return nil, user
}

func (this *User) Authorize(role string) (bool, error) {
	err, user := GetUserByUsername(this.Username)
	if err != nil {
		return false, err
	}

	fmt.Println("\n\nUser Role: ", user.Role)

	if role == "admin" {
		return true, nil
	}

	if strings.ToLower(user.Role) != strings.ToLower(role) {
		return false, errors.New("You're not authorized, because your role is not " + role)
	}
	return true, nil
}
