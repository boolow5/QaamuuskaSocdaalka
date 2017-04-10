package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/boolow5/QaamuuskaSocdaalka/models"
	"github.com/boolow5/bolow/bolOs"
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
	c.Data["Drafts"], _ = models.GetDrafts(page*itemsPerPage, itemsPerPage)
	c.Data["Categories"], _ = models.GetCategories()
	c.Data["Images"], _ = models.GetImages()
	c.Data["Users"], _ = models.GetUsers()
	c.Data["message"] = flash.Data
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
	fmt.Printf("Authentication:\nError: %v\nUser: %v", err, authenticatedUser)
	if err != nil {
		responseMessage["error"] = "user authentication error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	if authenticatedUser.Role == "" {
		responseMessage["error"] = "incorrect username or password"
		responseMessage["explation"] = "empty role"
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
		IsAdmin    bool   `json:"admin"`
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
			Email:      requestBody.Email,
		},
	}
	user.SetPassword(requestBody.Password)
	user.Role = "normal"
	if requestBody.IsAdmin {
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

func (this *AdminController) AddPost() {
	post := struct {
		Title           string `json:"title"`
		Content         string `json:"content"`
		CategoryId      string `json:"category"`
		FeaturedImageId string `json:"featured_image"`
		SaveAsDraft     bool   `json:"save_as_draft"`
	}{}

	responseMessage := map[string]interface{}{}
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&post)
	if err != nil {
		responseMessage["error"] = "post parsing error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	newPost := models.Post{
		Title:       post.Title,
		Content:     post.Content,
		SaveAsDraft: post.SaveAsDraft,
	}

	category_id, _ := strconv.Atoi(post.CategoryId)
	image_id, _ := strconv.Atoi(post.FeaturedImageId)
	newPost.Category = &models.Category{Id: category_id}
	newPost.FeaturedImage = &models.Image{Id: image_id}

	CurrentUser := this.GetSession("username")
	if CurrentUser == nil {
		responseMessage["error"] = "post-not-saved"
		responseMessage["explation"] = "invalid author"
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	_, newPost.Author = models.GetUserByUsername(CurrentUser.(string))

	saved := models.SaveItem(&newPost)
	if !saved {
		responseMessage["error"] = "post-not-saved"
		responseMessage["explation"] = ""
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	responseMessage["success"] = "added-post"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func (this *AdminController) AddCategory() {
	category := models.Category{}

	responseMessage := map[string]interface{}{}
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&category)
	if err != nil {
		responseMessage["error"] = "category parsing error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	saved := models.SaveItem(&category)
	if !saved {
		responseMessage["error"] = "category-not-saved"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	responseMessage["success"] = "added-category"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func (this *AdminController) AddImage() {
	fmt.Println("AddImage")
	flash := beego.NewFlash()
	image := models.Image{}

	imageTypes := []string{"image/jpeg", "image/png", "image/gif"}
	isImage := false

	// get the image file
	file, header, err := this.GetFile("file")
	if err != nil {
		flash.Error(i18n.Tr(this.Lang, "upload error") + "\n" + err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/bol-admin", 302)
		return
	}
	if file != nil {
		fileName := header.Filename
		fileType := header.Header.Get("Content-Type")
		fmt.Println("fileName", fileName)
		fmt.Println("fileType", fileType)

		for _, val := range imageTypes {
			if val == fileType {
				isImage = true
				break
			}
		}
	}

	if !isImage {
		flash.Error(i18n.Tr(this.Lang, "file not image"))
		flash.Store(&this.Controller)
		this.Redirect("/bol-admin", 302)
		return
	}
	// save the file
	uploadsDir := beego.AppConfig.String("uploads")
	shortName := header.Filename
	targetFile := uploadsDir + shortName
	// check a file with the same name exists
	for bolOs.FileExists(targetFile) {
		targetFile, shortName = bolOs.GenerateUniqueFileName(targetFile, "/")
	}

	// create file first
	outputFile, err := os.Create(targetFile)
	defer outputFile.Close()
	if err != nil {
		flash.Error(i18n.Tr(this.Lang, "creating file failed") + "\n" + err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/bol-admin", 302)
		return
	}

	if file != nil {
		fmt.Println("Target File:", targetFile)
		data, err := ioutil.ReadAll(file)
		if err != nil {
			flash.Error(i18n.Tr(this.Lang, "reading file failed") + "\n" + err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/bol-admin", 302)
			return
		}

		err = ioutil.WriteFile(targetFile, data, 644)
		// err = this.SaveToFile(fileName, targetFile)
		if err != nil {
			flash.Error(i18n.Tr(this.Lang, "saving file failed") + "\n" + err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/bol-admin", 302)
			return
		}
	}

	image.Title = this.GetString("title")
	image.Description = this.GetString("description")
	image.Url = "/static/uploads/" + shortName

	saved := models.SaveItem(&image)
	if !saved {
		flash.Error(i18n.Tr(this.Lang, "image not saved"))
		flash.Store(&this.Controller)
		this.Redirect("/bol-admin", 302)
		return
	}

	flash.Success(i18n.Tr(this.Lang, "image saved"))
	flash.Store(&this.Controller)
	this.Redirect("/bol-admin", 302)
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
