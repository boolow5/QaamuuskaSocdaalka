package routers

import (
	"github.com/boolow5/QaamuuskaSocdaalka/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
