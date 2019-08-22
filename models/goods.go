package models

import (
	"github.com/astaxie/beego/orm"
)

type Goods struct {
	Id int
	UserId int
	Name string
	ShowImgPath string
}

func init()  {
	orm.RegisterModel(new(Goods))
}

func GetAllGoods() []Goods {
	return nil
}
func AddGoods(goods Goods)Goods  {
	o:=orm.NewOrm()
	g:=new(Goods)
	g.Name=goods.Name
	g.UserId=goods.UserId
	g.ShowImgPath=goods.ShowImgPath
	_,err:=o.Insert(goods)
	if err != nil {
		panic(err.Error())
	}
	return *g
}
func GetGoodsByUser(id int)  {
	
}
