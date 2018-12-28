package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Merchant))
}

type Merchant struct {
	Id       int `orm:"pk;column(id);"`
	Name     string
	Password string
}
