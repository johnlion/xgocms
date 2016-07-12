package categorytag

import (
	"github.com/xgocms/controllers"
	"github.com/xgocms/extensions/services"
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

func