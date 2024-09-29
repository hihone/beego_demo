package main

import (
	"beego_demo/controllers"
	_ "beego_demo/db"
	_ "beego_demo/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
