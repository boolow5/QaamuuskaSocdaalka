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

func (c *MainController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	page, _ := strconv.Atoi(c.Ctx.Input.Query("page"))
	itemsPerPage, _ := strconv.Atoi(c.Ctx.Input.Query("ipp"))
	fmt.Println("Page:", page, "\tItems per Page:", itemsPerPage)
	if itemsPerPage == 0 {
		itemsPerPage = 20
	}
	c.Data["Posts"], _ = models.GetPosts(page*itemsPerPage, itemsPerPage)
	c.Data["message"] = flash
	SetTemplate("pages/index.tpl", &c.Controller)
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
	token := controller.GetSession("token")
	if username != nil && token != nil {
		// if found set the CurrentUser template variable and return username
		controller.Data["CurrentUser"] = username
		return username.(string)
	}

	// if session doesn't exist return empty string
	return ""
}
