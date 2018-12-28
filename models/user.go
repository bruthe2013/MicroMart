package models

import (
	"fmt"

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
func Login(username, password string) (int, int) {
	o := orm.NewOrm()
	fmt.Println(username)
	user := User{
		Username: username,
	}
	err := o.Read(&user, "username")
	if err == orm.ErrNoRows {
		//无此账号
		return 0, 0
	}
	if user.Password != password {
		return 1, 0
	}
	return 2, user.Id
}

//修改密码
func ModifyPassword(id int, password string) int {
	o := orm.NewOrm()
	user := User{Id: id}
	if o.Read(&user) == nil {
		user.Password = password
		if num, err := o.Update(&user); err == nil {
			fmt.Println("修改密码", num)
			return 1
		}
		return 1
	} else {
		return 2
	}
}
