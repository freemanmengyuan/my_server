package models

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type UserInfo struct {
	Id int64
	Name string
	Password string
}

func init() {
	sDbConf := GetDbConf()
	orm.Debug = true //是否开启调试模式
	orm.RegisterDataBase("default", "mysql", sDbConf, 30)
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
func GetUser(id int64, user *UserInfo){
	user.Id = id
	db.Read(&user)
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

/**
 * 校验用户
 */
func CheckUser(name string, user *UserInfo){
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("password").From("user_info").Where("name=?").Limit(1)
	sql := qb.String()
	db.Raw(sql, name).QueryRows(user)
}

/**
 * 获取数据库配置
 */
func GetDbConf() string {
	iniconf, _ := config.NewConfig("ini", "conf/app.conf")
	env := iniconf.String("runmode")
	dbhost := iniconf.String(env+"::dbhost")
	dbport := iniconf.String(env+"::dbport")
	dbuser := iniconf.String(env+"::dbuser")
	dbpassword := iniconf.String(env+"::dbpassword")
	str := dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/server?charset=utf8"
	return str
}


