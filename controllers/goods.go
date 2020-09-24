package controllers

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController)Get()  {
	c.Data["title"]="你好beego"
	c.TplName="goods.tpl"
}