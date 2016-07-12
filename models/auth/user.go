package auth


import (

)
import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Username string
	UserLevel int
}

type UserLoginForm struct {
	Username string                 `form:"username"`
	Password string                 `form:"password"`
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

func ( this *User ) CheckLogin( text string ) string{

	return "1234"
}






