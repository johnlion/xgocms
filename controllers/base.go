package controllers

import (
	"github.com/astaxie/beego"
	"github.com/johnlion/xgocms/extensions/services"
	"html/template"

	"github.com/johnlion/xgocms/models/utils"
	"github.com/johnlion/xgocms/setting"
	"github.com/beego/i18n"
	"time"
	"fmt"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type BaseController struct{
	beego.Controller
	services.Service
	themesAdmin string
	i18n.Locale

	Pagenums int
	Skip int
	Postdate string
	Uniqid int
	DB *mgo.Database

}

func ( this *BaseController ) SetThemesAdmin(){
	this.themesAdmin = beego.AppConfig.String("themesadmin")
}

func ( this *BaseController ) GetThemesAdmin() string{
	this.SetThemesAdmin()
	return this.themesAdmin
}


// XSRFFormHTML writes an input field contains xsrf token value.
func (this *BaseController) XSRFFormHTML() string {
	// return `<input type="hidden" name="_xsrf" ng-model="_xsrf" value="` +
	// this.XSRFToken() + `" />`
	return `<input type="hidden" name="_xsrf"  value="`+ this.XSRFToken() +`"   />`
}

func ( this *BaseController ) Prepare(){
	// conn database
	this.ConnDatabase()

	// user profile
	if  this.GetSession( "Authorized" ) == true{
		profile := map[interface{}]interface{}{
			"Authorized": this.GetSession( "Authorized" ) ,
			"Password": this.GetSession( "Password" ) ,
			"Uniqid": this.GetSession( "Uniqid" ) ,
			"UserLevel": this.GetSession( "UserLevel" ) ,
			"Username": this.GetSession( "Username" ) ,
		}
		this.Data["profile"] = profile
	}


	// page settings
	this.Pagenums = 10
	this.Postdate = time.Now().Format("2006-01-02 15:04:05")
	this.Data["xsrfdata"]= template.HTML(this.XSRFFormHTML())
	this.Data["page_title"] = "";
	this.Data["page_via_title"] = "";
	this.Data["IsLoginPage"] = false;

	// page layouts
	this.Layout = this.GetThemesAdmin() +"layout.html"
	this.LayoutSections = make( map[string]string )
	this.LayoutSections["Header"] = this.GetThemesAdmin() + "header.html"
	this.LayoutSections["Footer"] = this.GetThemesAdmin() + "footer.html"
	this.LayoutSections["Navbar"] = this.GetThemesAdmin() + "navbar.html"
	this.LayoutSections["Profile"] = this.GetThemesAdmin() + "profile.html"
	this.LayoutSections["Sidebar_menu"] = this.GetThemesAdmin() + "sidebar_menu.html"
	this.LayoutSections["MenuFooterButtons"] = this.GetThemesAdmin() + "menu_footer_buttons.html"
	this.LayoutSections["TopNavigation"] = this.GetThemesAdmin() + "top_navigation.html"

	this.setLang()

}

func (this *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}



func (this *BaseController) JsStorage(action, key string, values ...string) {
	value := action + ":::" + key
	if len(values) > 0 {
		value += ":::" + values[0]
	}
	this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/", nil, nil, false)
}

func (this *BaseController) setLangCookie(lang string) {
	this.Ctx.SetCookie("lang", lang, 60*60*24*365, "/", nil, nil, false)
}



// setLang sets site language version.
func (this *BaseController) setLang() bool {
	isNeedRedir := false
	hasCookie := false

	// get all lang names from i18n
	langs := setting.Langs

	// 1. Check URL arguments.
	lang := this.GetString("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	//// 3. check if isLogin then use user setting
	//if len(lang) == 0 && this.IsLogin {
	//	lang = i18n.GetLangByIndex(this.User.Lang)
	//}

	// 4. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Input.Header("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. DefaucurLang language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	// Save language information in cookies.
	if !hasCookie {
		this.setLangCookie(lang)
	}

	// Set language properties.
	this.Data["Lang"] = lang
	this.Data["Langs"] = langs

	this.Lang = lang

	return isNeedRedir
}

func ( this *BaseController ) ConnDatabase() {
	this.Service.MongoDBSession()
	this.MongoSession.SetMode( mgo.Strong ,true )
	this.DB = this.MongoSession.DB( beego.AppConfig.String("mongodb")  )

}

//　设置数据库查询记录开始下标
func ( this *BaseController ) SetSkip(){
	page, err := this.GetInt("p")
	beego.Debug( "Page",page )

	if err == nil{
		this.Skip = ( page - 1 ) * this.Pagenums
	}else{
		this.Skip  = 0
	}
}



// auto increment id
func ( this *BaseController ) GetSeq()int{
	counter := this.DB.C("counter")
	cid := "uniqid"

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Seq int }{}
	if _, err := counter.Find(bson.M{"_id": cid}).Apply(change, &doc); err != nil {
		panic(fmt.Errorf("get counter failed:", err.Error()))
	}
	return doc.Seq
}

func ( this *BaseController ) IsLogin(){
	if  this.GetSession( "Authorized" ) != true{
		this.Redirect( "/auth/login", 302 )
	}
}
