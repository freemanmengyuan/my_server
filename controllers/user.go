package controllers

import (
	"fmt"

	// 引入数据库驱动注册及初始化
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


type UserController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

type UserInfo struct {
	Id int64
	Username string
	Password string
}

func (c *UserController) Login() {
	u := User{}
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
 * model
 */
func (c *UserController) Regest() {
	orm.RegisterDataBase("default", "mysql", "root:$garo&&469312$@tcp(144.34.144.200:3306)/server?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	// 写入
	/*user := UserInfo{Username:"zhangsan", Password:"123456"}
	id,err := o.Insert(&user)
	if err == nil {
		c.Ctx.WriteString(fmt.Sprintf("insert user success %d ", id))
	}*/

	// 查询
	user := UserInfo{}
	user.Id = 1
	o.Read(&user)
	c.Ctx.WriteString(fmt.Sprintf("user info %v ", user))
}

/**
 * 原生sql
 */
func (c *UserController) Index() {
	orm.Debug = true //是否开启调试模式
	orm.RegisterDataBase("default", "mysql", "root:$garo&&469312$@tcp(144.34.144.200:3306)/server?charset=utf8", 30)
	//orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	/*var users []orm.Params
	o.Raw("select * from user_info").Values(&users)

	for _,v := range users{
		c.Ctx.WriteString(fmt.Sprintf("user %v ", v))
	}*/
	var users []UserInfo
	o.Raw("select * from user_info").QueryRows(&users)
	c.Ctx.WriteString(fmt.Sprintf("users %v", users))
}

/**
 * 查询构造器
 */
func (c *UserController) Info() {
	orm.Debug = true //是否开启调试模式
	orm.RegisterDataBase("default", "mysql", "root:$garo&&469312$@tcp(144.34.144.200:3306)/server?charset=utf8", 30)

	o := orm.NewOrm()

	var users []UserInfo
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("password").From("user_info").Where("username=?").Limit(1)
	sql := qb.String()
	o.Raw(sql, "zhangsan").QueryRows(&users)
	c.Ctx.WriteString(fmt.Sprintf("user info: %v", users))
}



