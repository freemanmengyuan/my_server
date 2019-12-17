package controllers

import (
	"fmt"
	"blog/models"
	"github.com/astaxie/beego"
)


type UsersController struct {
	beego.Controller
}

type UserForm struct {
	Username string
	Password string
}

/**
 * 登录校验
 */
func (c *UsersController) Login() {
	c.TplName = "login.tpl"
}


/**
 * 登录校验
 */
func (c *UsersController) LoginAction() {
	u := UserForm{}
	if err := c.ParseForm(&u); err != nil {
		//do something
		c.Ctx.WriteString("err")
	}

	c.Ctx.SetCookie("name", u.Username, 1000, "/")
	c.Data["username"] = u.Username
	c.TplName = "user.tpl"
	//c.Ctx.WriteString(u.Username + "您好啊 欢迎登录 Blog")
	//c.Ctx.WriteString( "<html><a href='/blog/edit'>点我</a></html>")
}

/**
 * 注册用户
 */
func (c *UsersController) RegestUser() {
	u := UserForm{}
	if err := c.ParseForm(&u); err != nil {
		//do something
		c.Ctx.WriteString("err")
	}
	user := models.UserInfo{Username:u.Username, Password:u.Password}
	id,err := models.AddUser(&user)
	if err == nil {
		c.Ctx.WriteString(fmt.Sprintf("user regest success %d ", id))
	}
}

func (c *UsersController) GetUsers() {

	var users []models.UserInfo
	models.GetUsers(&users)
	c.Data["users"] = users
	c.Data["len"] = len(users)
	c.TplName = "user.tpl"
}



