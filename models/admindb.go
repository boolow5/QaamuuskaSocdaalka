package models

import (
	"errors"
	"fmt"
	"log"
)

func GetCategories() ([]*Category, error) {
	categories := []*Category{}
	var err error
	_, err = o.QueryTable("category").OrderBy("-id").All(&categories)

	if err != nil {
		Verbose("GetCategories: %v", err)
		return categories, err
	}
	return categories, nil
}

func GetCountries() ([]*Country, error) {
	countries := []*Country{}
	var err error
	_, err = o.Raw("SELECT * FROM countries ORDER BY name").QueryRows(&countries) // .QueryTable("countries").OrderBy("-id").All(&countries)

	if err != nil {
		Verbose("GetCountries: %v", err)
		return countries, err
	}
	return countries, nil
}

func GetCities() ([]*City, error) {
	cities := []*City{}
	var err error
	_, err = o.Raw("SELECT * FROM Cities ORDER BY name").QueryRows(&cities) // .QueryTable("Cities").OrderBy("-id").All(&cities)

	if err != nil {
		Verbose("GetCities: %v", err)
		return cities, err
	}
	return cities, nil
}

func GetCountryById(id int) (Country, error) {
	country := Country{Id: id}
	var err error
	err = o.Raw("SELECT * FROM counties WHERE id = ?", id).QueryRow(&country)

	if err != nil {
		Verbose("GetCountryById: %v", err)
		return country, err
	}
	return country, nil
}

func GetCategoryById(id int) (Category, error) {
	category := Category{Id: id}
	var err error
	err = o.Raw("SELECT * FROM category WHERE id = ?", id).QueryRow(&category)

	if err != nil {
		Verbose("GetCategoryById: %v", err)
		return category, err
	}
	return category, nil
}

func GetImages() ([]*Image, error) {
	images := []*Image{}
	var err error
	_, err = o.QueryTable("images").OrderBy("-id").All(&images)

	if err != nil {
		Verbose("GetImages: %v", err)
		return images, err
	}
	return images, nil
}

// GetUserByUsername fetchs a user by its username field
func GetUsers() ([]*User, error) {
	users := []*User{}
	_, err := o.QueryTable("user").OrderBy("-id").All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

func UserExists(username string) bool {
	result := struct {
		Ids int `json:"ids" orm:"default(0)"`
	}{}

	err := o.Raw("SELECT COUNT(*) as ids FROM user WHERE username = ? ", username).QueryRow(&result)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if result.Ids == 0 {
		log.Printf("exists = %v\n", false)
		return false
	}
	return true
}

// CreateUser adds new user to the database
func CreateUser(item *User) (bool, error) {
	o.Begin()

	// check user if exists
	userExists := UserExists(item.Username)
	if userExists {
		return false, errors.New("Username already exists")
	}
	rows_affected, err := o.Insert(item.Profile)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No profile saved")
	}
	rows_affected, err = o.Insert(item)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No user inserted")
	}
	o.Commit()
	return true, nil
}

// UpdateItem modifies the non-empty fields in the newItem
func UpdateItem(oldItem, newItem MyModel) (bool, error) {
	o.Begin()
	err := o.Read(oldItem)
	if err != nil {
		return false, err
	}
	rows_affected, err := o.Update(newItem)
	if err != nil {
		fmt.Println("FailUpdate:", newItem)
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		fmt.Println("NotUpdatedItem:", newItem)
		o.Rollback()
		return false, errors.New("No rows affected")
	}
	err = o.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteItem removes items from the database and returns false and/or error if it fails
func DeleteItem(item MyModel) (bool, error) {
	o.Begin()
	rows_affected, err := o.Delete(item)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		return false, errors.New("0 Rows deleted")
	}
	o.Commit()
	return true, nil
}
