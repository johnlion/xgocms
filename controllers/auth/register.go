package auth

import (
	"github.com/johnlion/xgocms/controllers"
	"github.com/johnlion/xgocms/extensions/services"
)

type RegisterController struct{
	controllers.BaseController
	services.Service
}



func ( this *RegisterController ) Get(){
	this.Layout = ""
	this.TplName =  this.GetThemesAdmin() + "controllers/auth/register.html"
}

func ( this *RegisterController ) Register(){
	this.Layout = ""
	this.Data["IsLoginPage"] = true
	this.TplName =  this.GetThemesAdmin() + "controllers/auth/register.html"
}


