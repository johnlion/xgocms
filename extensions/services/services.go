package services

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)



type Service struct {
	MongoSession *mgo.Session
	UserId string
}

func ( this *Service ) MongoDBSession() *mgo.Session{

	if this.MongoSession == nil {
		var err error
		this.MongoSession , err = mgo.Dial( beego.AppConfig.String( "mongoaddr" ) )
		if err !=nil {
			panic(err)
		}
	}
	return this.MongoSession.Clone()
}

