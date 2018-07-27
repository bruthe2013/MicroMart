package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int `orm:"pk;column(id);"`
	Username string
	Password string
}

func AddUser(u User) User {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Id = u.Id
	user.Username = u.Username
	user.Password = u.Password
	_, err := o.Insert(user)
	if err != nil {
		panic(err.Error())
	}
	return *user
}

func GetUser(uid int) (u *User, err error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Id: uid}
	o.Read(&user)
	u = &user
	return u, err
}

//登录，0无账号，1密码错误 2成功
func Login(username, password string) int {
	o := orm.NewOrm()
	user := new(User)
	user.Username = username
	err := o.Read(user)
	if err == orm.ErrNoRows {
		//无此账号
		return 0
	}
	if user.Password != password {
		return 1
	}
	return 2
}
