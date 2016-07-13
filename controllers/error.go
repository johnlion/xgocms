package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func ( this *ErrorController ) Error401(){
	this.Data["content"] = "401　Unauthorized"
	this.TplName = "401.tpl"
}

func ( this *ErrorController ) Error404(){
	this.Data["content"] = "404 Not Found"
	this.TplName = "404.tpl"
}

func ( this *ErrorController ) Error500(){
	this.Data["content"] = "500 Internal Server Error"
	this.TplName = "500.tpl"
}

func ( this *ErrorController ) Error501(){
	this.Data["content"] = "501  Not Implemented"
	this.TplName = "501.tpl"
}

func ( this *ErrorController ) Error502(){
	this.Data["content"] = "Bad Gateway"
	this.TplName = "502.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}