package controllers

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/boolow5/QaamuuskaSocdaalka/models"
)

type MainController struct {
	BaseController
}

func GetPageItems(c *beego.Controller) {
	page, _ := strconv.Atoi(c.Ctx.Input.Query("page"))
	itemsPerPage, _ := strconv.Atoi(c.Ctx.Input.Query("ipp"))
	if itemsPerPage == 0 {
		itemsPerPage = 20
	}
	offset := page * itemsPerPage
	limit := itemsPerPage
	posts, _ := models.GetPosts(offset, limit)
	c.Data["MostPopularPosts"], _ = models.GetPopularPosts(3)
	latests, posts := latestPosts(posts, 3)
	c.Data["LatestPosts"] = latests
	c.Data["Posts"] = posts
	c.Data["Categories"], _ = models.GetCategories()
}

func (c *MainController) Get() {
	GetPageItems(&c.Controller)
	c.Data["Page"] = "home"
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["message"] = flash.Data
	SetTemplate("pages/index.tpl", &c.Controller)
}

func (c *MainController) GetPostDetail() {
	GetPageItems(&c.Controller)
	c.Data["Page"] = "datail"
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["message"] = flash.Data

	post_url := c.Ctx.Input.Param(":post_url")
	c.Data["Post"] = models.GetPostByUrl(post_url)
	SetTemplate("pages/detail.tpl", &c.Controller)
}

func latestPosts(posts []*models.Post, limit int) (latest, otherPosts []*models.Post) {
	// get the first 3 posts
	if limit < len(posts) {
		latest = posts[:limit]
	} else {
		latest = posts
	}
	fmt.Printf("Latests: %d\t\tOther posts: %d\n", len(latest), len(posts))
	// remove the 3 post from the original list
	otherPosts = posts[len(latest):]
	return latest, otherPosts
}

func SetTemplate(tplName string, controller *beego.Controller) {
	if len(controller.Layout) < 1 {
		controller.Layout = "layouts/base.tpl"
	}
	controller.TplName = tplName

	// set layout sections
	controller.LayoutSections = make(map[string]string)
	controller.LayoutSections["css"] = "partials/css.tpl"
	controller.LayoutSections["favicons"] = "partials/favicons.tpl"
	controller.LayoutSections["navbar"] = "partials/navbar.tpl"
	controller.LayoutSections["navbarAr"] = "partials/navbar-ar.tpl"
	controller.LayoutSections["featured"] = "partials/featured.tpl"
	controller.LayoutSections["sidebarRight"] = "partials/sidebar-right.tpl"
	controller.LayoutSections["sidebarLeft"] = "partials/sidebar-left.tpl"
	controller.LayoutSections["footer"] = "partials/footer.tpl"

	controller.Data["xsrf_token"] = controller.XSRFToken()
	controller.XSRFExpire = 7200
	controller.Data["xsrfdata"] = template.HTML(controller.XSRFFormHTML())

	// check for login
	controller.Data["LoggedIn"] = GetCurrentUser(controller) != ""
}

func GetCurrentUser(controller *beego.Controller) string {
	// check for session username
	username := controller.GetSession("username")
	role := controller.GetSession("role")
	if username != nil {
		// if found set the CurrentUser template variable and return username
		controller.Data["CurrentUser"] = username
		controller.Data["CurrentUserRole"] = role
		controller.Data["IsAdmin"] = role == "admin"
		return username.(string)
	}

	// if session doesn't exist return empty string
	return ""
}
