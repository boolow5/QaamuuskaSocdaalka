package routers

import (
	"github.com/astaxie/beego"
	"github.com/boolow5/QaamuuskaSocdaalka/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// font-facing urls
	beego.Router("/:post_url", &controllers.MainController{}, "get:GetPostDetail")

	// misc urls
	beego.Router("/translate/:word/:args", &controllers.MiscAPIController{}, "*:GetTranslation")

	// admin section urls
	beego.Router("/bol-admin", &controllers.AdminController{})
	beego.Router("/bol-admin/add/user", &controllers.AdminController{}, "post:AddUser")
	beego.Router("/bol-admin/add/category", &controllers.AdminController{}, "post:AddCategory")
	beego.Router("/bol-admin/update/category/:category_id", &controllers.AdminController{}, "post:UpdateCategory")
	beego.Router("/bol-admin/add/post", &controllers.AdminController{}, "post:AddPost")
	beego.Router("/bol-admin/update/post/:post_id", &controllers.AdminController{}, "post:UpdatePost")
	beego.Router("/bol-admin/add/image", &controllers.AdminController{}, "post:AddImage")
	beego.Router("/bol-admin/update/image/:image_id", &controllers.AdminController{}, "post:UpdateImage")
	beego.Router("/bol-admin/delete/image/:image_id", &controllers.AdminController{}, "get:DeleteImage")
	beego.Router("/bol-admin/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/bol-admin/logout", &controllers.AdminController{}, "*:Logout")

	beego.Router("/bol-admin/add/country", &controllers.AdminController{}, "get:GetWorldForm;post:AddCountry")
}
