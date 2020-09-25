package controllers

import (
	"WebProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1、获取请求数据
	userName:=c.Ctx.Input.Query("user")
	password:=c.Ctx.Input.Query("psd")
	//2、使用固定数据进行数据校验
	//admin 123456
	if userName!="admin" || password!="123546"{
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起,数据校验错误"))
		return
	}
	//校验正确的情况
   c.Ctx.ResponseWriter.Write([]byte("恭喜,数据校验成功。"))


	//c.Data["Website"] = "baidu.me"
	//c.Data["Email"] = "liulianwanfan@gmail.com"
	//c.TplName = "index.tpl"
}

//func (c *MainController)Post(){
//	//1.接受(解析)post请求参数
//	name:=c.Ctx.Request.FormValue("name")
//	age:=c.Ctx.Request.FormValue("age")
//	sex:=c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	//2.进行数据校验
//	if name!="admin"&&age!="18" &&sex!="男" {
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}



//该方法用于处理post类型的请求

func (c *MainController)Post(){
    //1、解析前端提交的json格式的数据
	var user models.User
	databytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据解析错误")
		return
	}
	err=json.Unmarshal(databytes,&user)
	if err != nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("用户",user.User)
	fmt.Println("缺口",user.Nick)
	fmt.Println("密码",user.Password)
	fmt.Println("生日",user.Birthday)
	fmt.Println("地址",user.Address)
	c.Ctx.WriteString("数据解析成功")
}
