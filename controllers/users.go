package controllers

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego"
)


type UsersController struct {
	beego.Controller
}

type UserForm struct {
	Name string
	Password string
}

/**
 * 登录页
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

	c.Ctx.SetCookie("name", u.Name, 1000, "/")
	c.Data["username"] = u.Name
	c.TplName = "user.tpl"
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
	user := models.UserInfo{Name:u.Name, Password:u.Password}
	id,err := models.AddUser(&user)
	if err == nil {
		c.Ctx.WriteString(fmt.Sprintf("user regest success %d ", id))
	}
}

/**
 * 获取所有用户
 */
func (c *UsersController) GetUsers() {

	var users []models.UserInfo
	models.GetUsers(&users)

	c.Ctx.Output.JSON(users, false, false)
}



