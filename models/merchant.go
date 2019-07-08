package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init() {
	orm.RegisterModel(new(Merchant))
}

type Merchant struct {
	Id       int `orm:"pk;column(id);"`
	Name     string
	Password string
}

func AddMerchant(u Merchant) Merchant {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Merchant)
	user.Id = u.Id
	user.Name = u.Name
	user.Password = u.Password
	_, err := o.Insert(user)
	if err != nil {
		panic(err.Error())
	}
	return *user
}

func GetMerchant(uid int) (u *User, err error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Id: uid}
	o.Read(&user)
	u = &user
	return u, err
}

//商家登录 0无账号，1密码错误 2成功
func MerchantLogin(name, password string) (int, int) {
	o := orm.NewOrm()
	user := Merchant{
		Name: name,
	}
	err := o.Read(&user, "name")
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
func MerchantModifyPassword(id int, password string) int {
	o := orm.NewOrm()
	user := Merchant{Id: id}
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
