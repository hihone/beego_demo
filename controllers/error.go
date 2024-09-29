package controllers

import beego "github.com/beego/beego/v2/server/web"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["title"] = "404"
	c.Data["content"] = "数据未找到"
	c.TplName = "errors/error.tpl"
}
func (c *ErrorController) Error401() {
	c.Data["title"] = "401"
	c.Data["content"] = "数据未找到"
	c.TplName = "errors/error.tpl"
}
func (c *ErrorController) Error501() {
	c.Data["title"] = "501"
	c.Data["content"] = "服务发送错误，请联系管理员"
	c.TplName = "errors/error.tpl"
}
