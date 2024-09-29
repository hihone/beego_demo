package controllers

import (
	"beego_demo/models"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/spf13/cast"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.hihone.cn"
	c.Data["Email"] = "yyanghf@gmail.com"
	c.TplName = "index/index.tpl"
}

func (c *MainController) AddPage() {
	c.Data["Title"] = "增加数据"
	c.TplName = "index/add.tpl"
}
func (c *MainController) Add() {
	book := AddRequest{}
	if err := c.ParseForm(&book); err != nil {
		panic("数据解析错误")
	}

	b := &models.Book{
		Title:   book.Title,
		Content: book.Content,
	}
	_, err := models.AddBook(b)
	if err != nil {
		panic("数据添加到数据库出错啦")
	}

	c.Abort("添加成功")

}

func (c *MainController) Info() {
	id := c.Ctx.Input.Query(":id")
	info, err := models.GetBookById(cast.ToInt(id))
	if err != nil {
		panic(err)
	}

	c.Data["title"] = "详情"
	c.Data["book"] = info
	c.TplName = "index/info.tpl"
}

func (c *MainController) Edit() {
	book := AddRequest{}
	if err := c.ParseForm(&book); err != nil {
		panic("数据解析错误")
	}

	info, err := models.GetBookById(cast.ToInt(book.Id))
	if err != nil {
		panic("信息不存在")
	}

	b := &models.Book{
		Id:      info.Id,
		Title:   book.Title,
		Content: book.Content,
	}
	err = models.UpdateBookById(b)
	if err != nil {
		panic("数据编辑保存到数据库出错啦")
	}

	c.Abort("保存成功")
}

type AddRequest struct {
	Id      string `form:"id"`
	Title   string `form:"title"`
	Content string `form:"content"`
}
