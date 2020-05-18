package routers

import (
	"beegodemo02/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// http://localhost:8080/api/123
	beego.Router("/api/:id", &controllers.ApiController{})

	// http://localhost:8080/cms_12.html
	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")
}
