package models

import (
	"gopkg.in/mgo.v2"
	"github.com/johnlion/xgocms/extensions/services"
	"github.com/astaxie/beego"
)

type BaseModule struct {
	DB *mgo.Database
	services.Service
}



func ( this *BaseModule ) ConnDatabase() {
	this.Service.MongoDBSession()
	this.MongoSession.SetMode( mgo.Strong ,true )
	this.DB = this.MongoSession.DB( beego.AppConfig.String("mongodb")  )

}