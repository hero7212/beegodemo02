package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

// http://localhost:8080/api/123
func (c *ApiController) Get() {
	// 获取动态路由的值
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口---" + id)
}
