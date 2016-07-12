package admin

import (
	"github.com/xgocms/controllers"
	"github.com/xgocms/extensions/services"
	"html/template"
)

type IndexController struct{
	controllers.BaseController
	services.Service


}

func ( this *IndexController) Get(){








	this.Data["xsrfdata"]= template.HTML(this.BaseController.XSRFFormHTML())
	this.TplName =  this.GetThemesAdmin() + "index.html"
}