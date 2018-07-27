package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (this *TestController)Get()  {
	this.Data["json"]="result:success"
	this.ServeJSON()

}