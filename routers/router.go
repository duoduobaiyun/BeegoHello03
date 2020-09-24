package routers

import (
	"WebProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
    beego.Router("/goods",&controllers.GoodsController{})
}
