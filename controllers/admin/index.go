package admin

import (
	"github.com/johnlion/xgocms/controllers"
	"github.com/johnlion/xgocms/extensions/services"
	"html/template"
	"github.com/astaxie/beego"
)

type IndexController struct{
	controllers.BaseController
	services.Service


}

func ( this *IndexController) Get(){
	this.IsLogin()




	this.Data["xsrfdata"]= template.HTML(this.BaseController.XSRFFormHTML())
	this.TplName =  this.GetThemesAdmin() + "index.html"
	beego.Debug( "x1 Username is " ,this.GetSession( "Username" ) )
	beego.Debug( "x2 Authorized is ", this.GetSession( "Authorized" ) )
}