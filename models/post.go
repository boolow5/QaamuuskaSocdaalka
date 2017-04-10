package models

import "time"

type Category struct {
	Id         int       `json:"id" orm:"auto"`
	Name       string    `json:"name" orm:"size(300)"`
	PostsCount int       `json:"posts_count" orm:"default(0)"`
	CreatedAt  time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

type Post struct {
	Id            int       `json:"id" orm:"auto"`
	Title         string    `json:"title" orm:"size(300)"`
	Content       string    `json:"content" orm:"size(5000)"`
	FeaturedImage *Image    `json:"featured_image" orm:"rel(fk)"`
	Category      *Category `json:"category" orm:"rel(fk)"`
	Views         int       `json:"views" orm:"default(0)"`
	Author        *User     `json:"author" orm:"rel(fk)"`
	SaveAsDraft   bool      `json:"save_as_draft"`
	Language      string    `json:"language" orm:"size(20)"`
	CreatedAt     time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

type Image struct {
	Id          int    `json:"id" orm:"auto"`
	Title       string `json:"title" orm:"size(300)"`
	Description string `json:"description" orm:"size(500)"`
	Url         string `json:"url"`
}

func (post *Post) TableName() string {
	return "posts"
}
func (this *Post) Valid() bool {
	return (len(this.Title) > 1 && len(this.Content) > 1 && this.Author.Id > 0)
}
func (this *Post) String() string {
	return this.Title
}

func (img *Image) TableName() string {
	return "images"
}
func (this *Image) Valid() bool {
	return (len(this.Url) > 1)
}
func (this *Image) String() string {
	return this.Title
}

// category
func (img *Category) TableName() string {
	return "category"
}
func (this *Category) Valid() bool {
	return (len(this.Name) > 1)
}
func (this *Category) String() string {
	return this.Name
}
