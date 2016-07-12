package auth

import (
	"github.com/xgocms/controllers"
	"github.com/xgocms/extensions/services"
	"github.com/astaxie/beego"
	"github.com/xgocms/models/auth"

)

type LoginController struct{
	controllers.BaseController
	services.Service
	auth.User
}



func ( this *LoginController ) Get_Login(){
	this.Layout = ""
	this.TplName =  this.GetThemesAdmin() + "controllers/auth/login.html"


	// no login

	// login

}

func ( this *LoginController ) Post_Login(){
	this.Layout = ""
	this.TplName =  this.GetThemesAdmin() + "controllers/auth/login.html"

	form := auth.UserLoginForm{}
	if err := this.ParseForm( &form ); err != nil{
		// do code
	}

	// crypt
	cryptStr := this.GetMD5Hash( beego.AppConfig.String("secret_key") + form.Username +  form.Password )
	beego.Debug( "cryptStr is " , cryptStr )

	// check login
	this.CheckLogin( cryptStr )


	auth.NewUserBy_Username_Password( form.Username, form.Password )
}

func ( this *LoginController) Logout(){

}


