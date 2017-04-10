package models

import "fmt"

func GetPosts(offset, limit int) ([]*Post, error) {
	news_items := []*Post{}
	var err error
	if offset == 0 && limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 AND show = 1 ORDER BY published_date DESC").QueryRows(&news_items)
	} else if offset == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 AND show = 1 ORDER BY published_date DESC LIMIT ?", limit).QueryRows(&news_items)
	} else if limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 AND show = 1 ORDER BY published_date DESC OFFSET ? ", offset).QueryRows(&news_items)
	} else {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 AND show = 1 ORDER BY published_date DESC LIMIT ? OFFSET ? ", limit, offset).QueryRows(&news_items)
	}

	if err != nil {
		Verbose("GetNewsItems: %v", err)
		return news_items, err
	}
	return news_items, nil
}

func GetAllPosts(offset, limit int) ([]*Post, error) {
	news_items := []*Post{}
	var err error
	if offset == 0 && limit == 0 {
		_, err = o.Raw("SELECT * FROM posts ORDER BY id DESC").QueryRows(&news_items)
	} else if offset == 0 {
		_, err = o.Raw("SELECT * FROM posts ORDER BY id DESC LIMIT ?", limit).QueryRows(&news_items)
	} else if limit == 0 {
		_, err = o.Raw("SELECT * FROM posts ORDER BY id DESC OFFSET ? ", offset).QueryRows(&news_items)
	} else {
		_, err = o.Raw("SELECT * FROM posts ORDER BY id DESC LIMIT ? OFFSET ? ", limit, offset).QueryRows(&news_items)
	}

	if err != nil {
		Verbose("GetNewsItems: %v", err)
		return news_items, err
	}
	return news_items, nil
}

func GetDrafts(offset, limit int) ([]*Post, error) {
	news_items := []*Post{}
	var err error
	if offset == 0 && limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 1 ORDER BY id DESC").QueryRows(&news_items)
	} else if offset == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 1 ORDER BY id DESC LIMIT ?", limit).QueryRows(&news_items)
	} else if limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 1 ORDER BY id DESC OFFSET ? ", offset).QueryRows(&news_items)
	} else {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 1 ORDER BY id DESC LIMIT ? OFFSET ? ", limit, offset).QueryRows(&news_items)
	}

	if err != nil {
		Verbose("GetNewsItems: %v", err)
		return news_items, err
	}
	return news_items, nil
}

// GetUserByUsername fetchs a user by its username field
func GetUserByUsername(username string) (error, *User) {
	user := User{Username: username}
	err := o.Read(&user, "username")
	if err != nil {
		return err, &user
	}
	return nil, &user
}

func SaveItem(m MyModel) bool {
	if !m.Valid() {
		return false
	}
	i, err := o.Insert(m)
	if i < 1 || err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func GetItemById(id int, out MyModel) {
	if typeOf(out) == "models.User" {
		user := User{Id: id}
		err := o.Read(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &user
			return
		}
	} else if typeOf(out) == "models.Post" {
		post := Post{Id: id}
		err := o.Read(&post)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &post
			return
		}
	} else if typeOf(out) == "models.Image" {
		img := Image{Id: id}
		err := o.Read(&img)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &img
			return
		}
	}
}

func typeOf(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
