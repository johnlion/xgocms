package controllers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"

	"fmt"
)

type spiderRecord struct{
	Url string
	AddTime string
	UpdateTime string
}


type ListController struct {
	BaseController

}

func ( this *ListController ) Get(){
	fmt.Println( beego.AppConfig.String("mongoaddr") )
	monogosession := this.MongoDBSession()
	defer monogosession.Close()
	monogosession.SetMode( mgo.Strong ,true )
	db := monogosession.DB( "www_750xs_com")
	var result spiderRecord
	db.C( "spider" ).Find( bson.M{"url":"http://www.750xs.com/10769.html"} ).One( &result )


	this.Ctx.WriteString( result.Url )
}