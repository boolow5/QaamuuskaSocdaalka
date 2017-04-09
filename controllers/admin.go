package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/boolow5/QaamuuskaSocdaalka/models"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	page, _ := strconv.Atoi(c.Ctx.Input.Query("page"))
	itemsPerPage, _ := strconv.Atoi(c.Ctx.Input.Query("ipp"))
	fmt.Println("Page:", page, "\tItems per Page:", itemsPerPage)
	if itemsPerPage == 0 {
		itemsPerPage = 20
	}
	c.Data["Posts"], _ = models.GetPosts(page*itemsPerPage, itemsPerPage)
	c.Data["Categories"], _ = models.GetCategories()
	c.Data["Users"], _ = models.GetUsers()
	c.Data["message"] = flash
	SetAdminTemplate("admin/index.tpl", &c.Controller)
}

func (this *AdminController) Login() {
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	responseMessage := map[string]interface{}{}
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&requestBody)
	if err != nil {
		responseMessage["error"] = "iweydi-response-parsing-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	user := models.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	}
	err, authenticatedUser := user.Authenticate()
	if err != nil {
		responseMessage["error"] = "user authentication error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	if authenticatedUser.Role == "" {
		responseMessage["error"] = "incorrect username or password"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	fullname := authenticatedUser.Profile.FirstName
	if len(authenticatedUser.Profile.MiddleName) > 1 {
		fullname += " " + authenticatedUser.Profile.MiddleName
	}
	if len(authenticatedUser.Profile.LastName) > 1 {
		fullname += " " + authenticatedUser.Profile.LastName
	}

	this.SetSession("username", authenticatedUser.Username)
	this.SetSession("fullname", fullname)
	this.SetSession("role", authenticatedUser.Role)

	responseMessage["success"] = "login-success"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func (this *AdminController) Logout() {
	this.DestroySession()
	responseMessage := map[string]interface{}{}
	responseMessage["success"] = "logout-success"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func (this *AdminController) AddUser() {
	requestBody := struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
		IsAdmin    string `json:"admin"`
	}{}

	responseMessage := map[string]interface{}{}
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&requestBody)
	if err != nil {
		responseMessage["error"] = "iweydi-response-parsing-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	user := models.User{
		Username: requestBody.Username,
		Profile: &models.Profile{
			FirstName:  requestBody.FirstName,
			MiddleName: requestBody.MiddleName,
			LastName:   requestBody.LastName,
		},
	}
	user.SetPassword(requestBody.Password)
	user.Role = "normal"
	is_admin, _ := strconv.ParseBool(requestBody.IsAdmin)
	if is_admin {
		user.Role = "admin"
	}

	saved, err := user.Add()
	if err != nil {
		responseMessage["error"] = "user-saving-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	if !saved {
		responseMessage["error"] = "user-not-saved"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	responseMessage["success"] = "added-user"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func SetAdminTemplate(tplName string, controller *beego.Controller) {
	if len(controller.Layout) < 1 {
		controller.Layout = "layouts/admin.tpl"
	}
	controller.TplName = tplName

	// set layout sections
	controller.LayoutSections = make(map[string]string)
	controller.LayoutSections["css"] = "partials/css.tpl"
	controller.LayoutSections["favicons"] = "partials/favicons.tpl"
	controller.LayoutSections["navbar"] = "admin/partials/navbar.tpl"
	controller.LayoutSections["featured"] = "partials/featured.tpl"
	controller.LayoutSections["sidebarRight"] = "admin/partials/sidebar-right.tpl"
	controller.LayoutSections["sidebarLeft"] = "partials/sidebar-left.tpl"
	controller.LayoutSections["footer"] = "partials/footer.tpl"

	controller.Data["xsrf_token"] = controller.XSRFToken()
	controller.XSRFExpire = 7200
	controller.Data["xsrfdata"] = template.HTML(controller.XSRFFormHTML())

	// check for login
	controller.Data["LoggedIn"] = GetCurrentUser(controller) != ""
}
