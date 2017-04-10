package routers

import (
	"github.com/astaxie/beego"
	"github.com/boolow5/QaamuuskaSocdaalka/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/translate/:word/:args", &controllers.MiscAPIController{}, "*:GetTranslation")
	beego.Router("/bol-admin", &controllers.AdminController{})
	beego.Router("/bol-admin/add/user", &controllers.AdminController{}, "post:AddUser")
	beego.Router("/bol-admin/add/category", &controllers.AdminController{}, "post:AddCategory")
	beego.Router("/bol-admin/add/post", &controllers.AdminController{}, "post:AddPost")
	beego.Router("/bol-admin/add/image", &controllers.AdminController{}, "post:AddImage")
	beego.Router("/bol-admin/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/bol-admin/logout", &controllers.AdminController{}, "*:Logout")
}
