package auth

import (
	"github.com/johnlion/xgocms/controllers"
	"github.com/johnlion/xgocms/extensions/services"
	"github.com/johnlion/xgocms/models/auth"
	"github.com/johnlion/xgocms/models/utils"
	"github.com/astaxie/beego"
)

type LoginController struct{
	controllers.BaseController
	services.Service
	auth.User
}



func ( this *LoginController ) Get_Login(){
	if  this.GetSession( "Authorized" ) == true{
		this.Redirect( "/admin/index", 302 )
	}

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



	// login
	userEntity := this.User.Login( form  )
	if userEntity.Username ==""{
		//do code
		this.Data["json"] = map[string]interface{}{ "errorCode": 0, "msg":"用户名称或密码错误!" }
		this.ServeJSON()
	}else{
		this.SetSession( "Authorized", true )
		this.SetSession( "Username", userEntity.Username )
		this.SetSession( "UserLevel", this.GetMD5Hash( userEntity.Username + userEntity.Password   ) )
		this.SetSession( "Uniqid", this.GetMD5Hash(  utils.ToStr( userEntity.Uniqid ) + userEntity.Username ) )
		beego.Debug( "debug ...", this.GetSession( "Username" ) )

		this.Data["json"] = map[string]interface{}{ "errorCode": 0, "msg":"登录成功!请稍等登录中 . . .　" }
		this.ServeJSON()
	}




}

func ( this *LoginController) Logout(){
	this.DestroySession()
	this.Redirect( "/auth/login",302 )
}


