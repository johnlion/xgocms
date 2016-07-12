package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xgocms/routers"

	"github.com/xgocms/setting"

	"github.com/xgocms/controllers/auth"

	"github.com/xgocms/controllers/admin/node"
	
)

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize(){
	//setting.Loadconfig()

	//if err := utils.InitSphinxPools(); err != nil {
	//	beego.Error(fmt.Sprint("sphinx init pool", err))
	//}

	//SocialAuth
	//setting.SocialAuth = social.NewSocial("/login/", auth.SocialAuther)
	//setting.SocialAuth.ConnectSuccessURL = "/settings/profile"
	//setting.SocialAuth.ConnectFailedURL = "/settings/profile"
	//setting.SocialAuth.ConnectRegisterURL = "/register/connect"
	//setting.SocialAuth.LoginURL = "/login"

}

func main() {


	beego.SetLogFuncCall(true)

	beego.Info( "AppPath", beego.AppPath )

	if setting.IsProMode {
		beego.Info("Product mode enabled")
	} else {
		beego.Info("Develment mode enabled")
	}

	beego.Info(beego.BConfig.AppName, setting.APP_VER, setting.AppUrl)
	if !setting.IsProMode {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	// Register routers.
	// [example 1]
	/*
	//beego.Router("/mytest", &controllers.MyTestController{})
	//beego.Router("/list", &controllers.ListController{})

	// admin
	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	//beego.Router("/admin/node/add", &node.AddController{})
	//beego.Router("/admin/node/list:page([0-9]+)", &node.ListController{})
	*/
	// [example 2]
	login := new ( auth.LoginController  )
	beego.Router( "/auth/login", login, "get:Get_Login;post:Post_Login"  )
	beego.Router("/logout", login, "get:Logout")

	register := new ( auth.RegisterController  )
	beego.Router("/register", register, "get:Get" )

	admin_node := new( node.NodeController )
	beego.Router( "/admin/node/add", admin_node, "get:Get_add;post:Post_add"  )
	beego.Router( "/admin/node/edit/:Uniqid", admin_node, "get:Get_edit;post:Post_edit"  )
	beego.Router( "/admin/node/list", admin_node, "get:Get_list"  )



	beego.Run()
}
