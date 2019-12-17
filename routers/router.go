package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // 博客展示
	beego.Router("/blog/index", &controllers.BlogController{}, "get:List")
	beego.Router("/blog/info", &controllers.BlogController{}, "get:Info")

	// 博客管理
	beego.Router("/blog/edit", &controllers.BlogController{}, "get:Edit")
	beego.Router("/blog/add", &controllers.BlogController{}, "post:Update")
	beego.Router("/blog/del", &controllers.BlogController{}, "delete:Del")

    // 用户相关
	beego.Router("/user/login", &controllers.UserController{}, "Post:Login")
	beego.Router("/user/regest", &controllers.UserController{}, "Get:Regest")
	beego.Router("/user/index", &controllers.UserController{}, "Get:Index")
	beego.Router("/user/info", &controllers.UserController{}, "Get:Info")

	// 用户相关
	beego.Router("/users/login", &controllers.UsersController{}, "Get:Login")
	beego.Router("/users/login", &controllers.UsersController{}, "Post:LoginAction")
	beego.Router("/users/regest", &controllers.UsersController{}, "Get:RegestUser")

}
