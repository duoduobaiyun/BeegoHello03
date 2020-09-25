package routers

import (
	"WebProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/goods",&controllers.GoodsController{})
}
