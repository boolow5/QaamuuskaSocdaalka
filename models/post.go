package models

import "time"

type Post struct {
	Id            int       `json:"id" orm:"auto"`
	Title         string    `json:"title" orm:"size(300)"`
	Content       string    `json:"content" orm:"size(5000)"`
	FeaturedImage *Image    `json:"featured_image" orm:"rel(fk)"`
	Views         int       `json:"views" orm:"default(0)"`
	Author        *User     `json:"author" orm:"rel(fk)"`
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
	return (len(this.Title) > 1 && len(this.Content) > 1 && this.Author.Id > 0 && this.FeaturedImage.Url != "")
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
