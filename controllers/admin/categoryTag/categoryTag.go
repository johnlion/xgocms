package categorytag

import (
	"github.com/johnlion/xgocms/controllers"
	"github.com/johnlion/xgocms/extensions/services"
)

type CategoryTag struct{
	Title string
	Hot int
	AddTime string
	UpdateTime string
	Sort int

}

type Form struct{
	Title string            `form:"title"`
	Sort int                `form:"sort"`
}

type CategoryTagController struct {
	controllers.BaseController
	services.Service
}

