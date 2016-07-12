package models

import (
	"gopkg.in/mgo.v2"
	"github.com/xgocms/extensions/services"
	"github.com/astaxie/beego"
)

type BaseModule struct {
	DB *mgo.Database
}

func ( this *BaseModule ) Conndata() *mgo.Database{
	monogosession := services.Service.MongoDBSession()
	monogosession.SetMode( mgo.Strong ,true )
	this.DB = monogosession.DB( beego.AppConfig.String("mongodb")  )
	return this.DB
}
