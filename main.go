package main

import (
	_ "WebProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

//今天遇到的错误，没有正确理解router包里面的 beego.Router的作用,它&controllers.后面接的是结构体,记住结构体!!!!!,
//现在跟的是resisterController里面的结构体!!!!!!!,主要是router包跟controllers包的关系与作用,当然还是跟结构体有关,但是
//还是得多多理解包于其他包的关系
