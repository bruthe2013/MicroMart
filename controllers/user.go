package controllers

import (
	"fmt"
	"../models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["json"] = "test"
	c.ServeJSON()
}
func (this *UserController) Add() {

	var u models.User
	u.Id, _ = this.GetInt("id")
	u.Username = this.GetString("username")
	u.Password = this.GetString("password")

	models.AddUser(u)
	user := models.AddUser(u)
	this.Data["json"] = user
	this.ServeJSON()
}

func (c *UserController) Getinfo() {
	id, _ := c.GetInt("id")
	fmt.Println(id)
	us, err := models.GetUser(id)
	if err != nil {
		c.Data["json"] = "user not exist"
	}
	c.Data["json"] = us
	c.ServeJSON()
}

func (this *UserController) Login() {
	p := this.GetString("password")
	n := this.GetString("username")
	switch models.Login(n, p) {
	case 0:
		this.Data["json"] = "user not exist"
		break
	case 1:
		this.Data["json"] = "password err"
		break
	case 2:
		this.Data["json"] = "login success"
		break
	}
	this.ServeJSON()
}
