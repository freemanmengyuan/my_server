package controllers

import (
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

type Article struct {
	Title string
	Content string
}

func (c *BlogController) List() {
	c.Ctx.WriteString("列表")
}

func (c *BlogController) Info() {
	id := c.GetString("id")
	c.Ctx.WriteString(id)
}

func (c *BlogController) Edit() {
	name := c.Ctx.GetCookie("name")
	if name == "" {
		c.TplName = "login.tpl"
	} else {
		c.Data["title"] = name + "的博客编辑页"
		c.TplName = "form.tpl"
	}
}

func (c *BlogController) Update() {
	u := Article{}
	if err := c.ParseForm(&u); err != nil {
		//do something
		c.Ctx.WriteString("err")
	}
	c.Ctx.WriteString(u.Title + u.Content)
}

func (c *BlogController) Del() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
