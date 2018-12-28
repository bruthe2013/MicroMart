// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"MicroMart/controllers"
	"fmt"
	"unicode/utf8"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/img/:name", &controllers.UserController{}, "get:DownLoadImg")
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/getinfo", &controllers.UserController{}, "get:Getinfo")
	beego.Router("/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/user/modify", &controllers.UserController{}, "post:ModifyPassword")
	fmt.Println("注册路由成功")

}
