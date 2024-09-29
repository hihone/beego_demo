package routers

import (
	"beego_demo/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

func init() {
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace(
		"/api_demo",
		beego.NSRouter("/add", &controllers.MainController{}, "get:AddPage"),
		beego.NSRouter("/add", &controllers.MainController{}, "post:Add"),
		beego.NSRouter("/info/?:id", &controllers.MainController{}, "get:Info"),
		beego.NSRouter("/edit/:id", &controllers.MainController{}, "put:Edit"),
	)
	beego.AddNamespace(ns)
}
