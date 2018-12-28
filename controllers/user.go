package controllers

import (
	"MicroMart/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type UserController struct {
	beego.Controller
}
type UserResp struct {
	StatusCode int
	Msg        string
}
type LoginResp struct {
	UserResp
	LoginState int
	Token      string
	Id         int
}
type ModifyPassword struct {
	UserResp
	State int
}

const (
	SecretKey = "MicroMart"
)

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
	//返回0用户不存在，1密码错误，2成功
	p := this.GetString("password")
	n := this.GetString("username")
	lr := LoginResp{}
	lr.StatusCode = 200
	lr.Msg = "success"
	if lr.LoginState, lr.Id = models.Login(n, p); lr.LoginState == 2 {
		lr.Token, _ = getToken()
	}
	this.Data["json"] = lr
	this.ServeJSON()
}
func getToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}
func (this *UserController) ModifyPassword() {
	uid, _ := this.GetInt("uid")
	p := this.GetString("password")
	status := models.ModifyPassword(uid, p)
	mr := ModifyPassword{State: status}
	mr.StatusCode = 200
	mr.Msg = "OK"
	this.Data["json"] = mr
	this.ServeJSON()
}

//下载图片
func (this *UserController) DownLoadImg() {
	this.Ctx.Output.Download("D:/file_sys/"+this.Ctx.Input.Param(":name"), "test.png")
}
