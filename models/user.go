package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type UserInfo struct {
	Id int64
	Username string
	Password string
}

func init() {
	orm.Debug = true //是否开启调试模式
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/server?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/server?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

/**
 * 添加用户
 */
func AddUser(user *UserInfo) (int64, error){
	id,err := db.Insert(user)
	return id,err
}

/**
 * 获取用户信息
 */
func GetUser(user *UserInfo) (int64, error){
	id,err := db.Insert(user)
	return id,err
}

/**
 * 获取所有用户信息
 */
func GetUsers(users *[]UserInfo){
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}




