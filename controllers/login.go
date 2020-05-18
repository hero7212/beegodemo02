package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//
// 模板
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {

	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)

	// 执行跳转

	// c.Redirect("/", 302)
	// c.Redirect("/cms_1234.html", 302)

	c.Ctx.Redirect(302, "/")
}
