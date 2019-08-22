package controllers

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
)

type MerchantController struct {
	beego.Controller
}
type MerchantResp struct {
	StatusCode int
	Msg        string
}
type MerchantLoginResp struct {
	UserResp
	LoginState int
	Token      string
	Id         int
}
type MerchantModifyPassword struct {
	UserResp
	State int
}
func (this *MerchantController) Add() {
	var u models.Merchant
	u.Id, _ = this.GetInt("id")
	u.Name = this.GetString("name")
	u.Password = this.GetString("password")
	models.AddMerchant(u)
	user := models.AddMerchant(u)
	this.Data["json"] = user
	this.ServeJSON()
}

func (c *MerchantController) Getinfo() {
	id, _ := c.GetInt("id")
	fmt.Println(id)
	us, err := models.GetUser(id)
	if err != nil {
		c.Data["json"] = "user not exist"
	}
	c.Data["json"] = us
	c.ServeJSON()
}
func (this *MerchantController) Login() {
	//返回0用户不存在，1密码错误，2成功
	p := this.GetString("password")
	n := this.GetString("name")
	lr := MerchantLoginResp{}
	lr.StatusCode = 200
	lr.Msg = "success"
	if lr.LoginState, lr.Id = models.Login(n, p); lr.LoginState == 2 {
		lr.Token, _ = getToken()
	}
	this.Data["json"] = lr
	this.ServeJSON()
}
func (this *MerchantController) ModifyPassword() {
	uid, _ := this.GetInt("uid")
	p := this.GetString("password")
	status := models.MerchantModifyPassword(uid, p)
	mr := MerchantModifyPassword{State: status}
	mr.StatusCode = 200
	mr.Msg = "OK"
	this.Data["json"] = mr
	this.ServeJSON()
}
