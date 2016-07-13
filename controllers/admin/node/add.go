package node

import (
	"gopkg.in/mgo.v2/bson"
)

func ( this *NodeController ) Get_add(){
	this.IsLogin()
	this.Data["page_title"] = "Node";
	this.Data["page_via_title"] = "Add Form";
	this.TplName =  this.GetThemesAdmin() + "controllers/node/add.html"
}

func ( this *NodeController ) Post_add(){
	this.IsLogin()
	form := Form{}
	if err := this.ParseForm( &form ); err != nil{
		// do code

	}

	var result Node
	err := this.DB.C( "node" ).Find(bson.M{"title": form.Title  }).One(&result)
	if err != nil{
		// this.Ctx.WriteString( "table node no  values" )
	}
	if  result.Title == "" {
		uniqid := this.GetSeq()

		this.DB.C( "node" ).Insert(&Node{ "News" , form.Title, form.Content , 0  , this.Postdate, "" ,form.Description, form.Keywords, form.Author, uniqid  })
		this.Data["json"] = map[string]interface{}{ "errorCode": 0, "msg":"添加成功!" }
		this.ServeJSON()
		return
	}else{
		this.Data["json"] = map[string]interface{}{ "errorCode": 1, "msg":  "title " + form.Title + "had existed" }
		this.ServeJSON()
		return
	}
}


