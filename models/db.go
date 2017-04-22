package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

func GetPopularPosts(limit int) (popularPosts []*Post, err error) {
	_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 ORDER BY views DESC LIMIT ?", limit).QueryRows(&popularPosts) // o.QueryTable("posts").Filter("save_as_draft", 0).OrderBy("views").Limit(limit).All(&popularPosts)
	for i := 0; i < len(popularPosts); i++ {
		o.QueryTable("images").Filter("id", popularPosts[i].FeaturedImage.Id).One(popularPosts[i].FeaturedImage)
	}
	return popularPosts, err
}

func GetPostById(post_id int) (post Post) {
	// get the post
	err := o.Raw("SELECT * FROM posts WHERE id = ? AND save_as_draft = 0", post_id).QueryRow(&post)
	if err != nil {
		fmt.Println(err)
		return post
	}
	o.Raw("SELECT * FROM images WHERE id = ?", post.FeaturedImage.Id).QueryRow(post.FeaturedImage)

	post.Views += 1
	o.Raw("UPDATE posts SET views = ? WHERE id = ?", post.Views, post_id).Exec()
	return post
}

func GetPostByUrl(post_url string) (post Post) {
	// get the post
	err := o.Raw("SELECT * FROM posts WHERE url = ? AND save_as_draft = 0", post_url).QueryRow(&post)
	if err != nil {
		fmt.Println(err)
		return post
	}
	o.Raw("SELECT * FROM images WHERE id = ?", post.FeaturedImage.Id).QueryRow(post.FeaturedImage)

	post.Views += 1
	o.Raw("UPDATE posts SET views = ? WHERE id = ?", post.Views, post.Id).Exec()
	return post
}

func GetPosts(offset, limit int) ([]*Post, error) {
	news_items := []*Post{}
	var err error
	if offset == 0 && limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 ORDER BY published_date DESC").QueryRows(&news_items)
		// o.QueryTable("posts").Filter("save_as_draft", 0).OrderBy("-published_date").Offset(offset).Limit(limit).All(&news_items)
	} else if offset == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 ORDER BY published_date DESC LIMIT ?", limit).QueryRows(&news_items)
		// o.QueryTable("posts").Filter("save_as_draft", 0).OrderBy("-published_date").Limit(limit).All(&news_items)
	} else if limit == 0 {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 ORDER BY published_date DESC OFFSET ? ", offset).QueryRows(&news_items)
		// o.QueryTable("posts").Filter("save_as_draft", 0).OrderBy("-published_date").Offset(offset).All(&news_items)
	} else {
		_, err = o.Raw("SELECT * FROM posts WHERE save_as_draft = 0 ORDER BY published_date DESC LIMIT ? OFFSET ? ", limit, offset).QueryRows(&news_items)
		// o.QueryTable("posts").Filter("save_as_draft", 0).OrderBy("-published_date").All(&news_items)
	}

	if err != nil {
		Verbose("GetNewsItems: %v", err)
		return news_items, err
	}

	for i := 0; i < len(news_items); i++ {
		o.QueryTable("images").Filter("id", news_items[i].FeaturedImage.Id).One(news_items[i].FeaturedImage)
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
	if id < 1 {
		return
	}
	out.SetId(id)
	err := o.Read(out)
	if err != nil {
		fmt.Println(err)
	}
	/*if typeOf(out) == "models.User" {
		user := User{Id: id}
		err := o.Raw("SELECT * FROM ? WHERE id = ?", out.TableName(), id).QueryRow(&user) // QueryTable(out.TableName()).Filter("id = ?", id).One(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &user
			return
		}
	} else if typeOf(out) == "models.Post" {
		post := Post{Id: id}
		err := o.Raw("SELECT * FROM ? WHERE id = ?", out.TableName(), id).QueryRow(&post) // o.Read(&post)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &post
			return
		}
	} else if typeOf(out) == "models.Image" {
		img := Image{Id: id}
		err := o.Raw("SELECT * FROM ? WHERE id = ?", out.TableName(), id).QueryRow(&img) // o.Read(&img)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &img
			return
		}
	} else if strings.HasSuffix(typeOf(out), "models.Category") {
		category := Category{Id: id}
		err := o.Raw("SELECT * FROM ? WHERE id = ?", out.TableName(), id).QueryRow(&category) // o.Read(&category)
		if err != nil {
			fmt.Println(err)
		} else {
			out = &category
			return
		}
	}*/
}

func UpdateAllCategoriesPostCount() {
	categories := []Category{}
	// get all categories
	o.QueryTable("category").All(&categories)

	// iterate through each and update it
	for i := 0; i < len(categories); i++ {
		o.Raw("UPDATE category SET posts_count = (SELECT COUNT(*) FROM posts WHERE category_id = ?) WHERE id = ?", categories[i].Id, categories[i].Id).Exec()
	}
}

func UpdateCategoryPostCount(category_id int) {
	category := Category{Id: category_id}
	counterValues := []orm.Params{}
	o.Raw("SELECT COUNT(*) as posts FROM posts WHERE category_id = ?", category_id).Values(&counterValues)
	counter := 0
	if counterValues[0] != nil {
		counter, _ = strconv.Atoi(counterValues[0]["posts"].(string))
	}
	fmt.Println("Counter:", counter, "\nCounterValues:", counterValues[0])
	category.PostsCount = int(counter)
	result, err := o.Raw("UPDATE category SET posts_count = ? WHERE id = ?", counter, category_id).Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result:", result)

}

func typeOf(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
