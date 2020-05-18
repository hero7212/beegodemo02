package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

// 路由伪静态
func (c *CmsController) Get() {

	cmsId := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms详情---" + cmsId)
}
