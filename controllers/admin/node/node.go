package node

import (
	"github.com/johnlion/xgocms/controllers"
	"github.com/johnlion/xgocms/extensions/services"

)

type Node struct{
	CategoryTag string
	Title string
	Content string
	Hot int
	AddTime string
	UpdateTime string
	Description string
	Keywords string
	Author string
	Uniqid int


}

type Form struct {
	Title string            `form:"title"`
	Content string          `form:"content"`
	Description string      `form:"description"`
	Keywords string         `form:"keywords"`
	Author string           `form:"author"`
	CategoryTag string      `form:"categorytag"`

}

type EditForm struct{
	Title string            `form:"title"`
	Content string          `form:"content"`
	Description string      `form:"description"`
	Keywords string         `form:"keywords"`
	Author string           `form:"author"`
	CategoryTag string      `form:"categorytag"`
	Hot int                 `form:"hot"`
	AddTime string          `form:"addtime"`
	Uniqid string           `form:"uniqid"`
}


type spiderRecord struct{
	Url string
	AddTime string
	UpdateTime string
}


type NodeController struct{
	controllers.BaseController
	services.Service
}







