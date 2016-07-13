package auth


import (

)
import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
	"github.com/johnlion/xgocms/models"
	"gopkg.in/mgo.v2/bson"

)


type UserEntity struct {
	Username string
	UserLevel int
	Password string
	Uniqid int

}

type UserLoginForm struct {
	Username string                 `form:"username"`
	Password string                 `form:"password"`
	Logintype string                `form:"logintype"`
}


type User struct {
	models.BaseModule
	UserEntity


}


func NewUser () {

}

//
func NewUserBy_Username_Password( username string, password string ){
	beego.Debug( "username is ", username  )
	beego.Debug( "password is ", password )
}

func ( this *User ) GetUsername() string{
	return this.Username
}
func ( this *User ) GetUserLevel() int{
	return this.UserLevel
}

func ( this *User ) GetMD5Hash( text string ) string{
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ( this *User ) Login( form interface{} )  UserEntity{
	if inst, ok := form.(UserLoginForm); ok  {
		switch inst.Logintype {
		case "direct":
			return this.Direct( inst.Username, inst.Password )
			break
		case "third":
			return UserEntity{}
			break
		default:
			return UserEntity{}
		}
	}
	return UserEntity{}
}

func ( this *User )  Direct( Username string,  Password string ) UserEntity{
	this.Username = Username
	this.Password = this.GetMD5Hash(  beego.AppConfig.String("secret_key") +Username +   Password )
	this.ConnDatabase()

	var result UserEntity
	this.DB.C("user").Find( bson.M{"username": this.Username, "password": this.Password}).One( &result )
	return result
}




