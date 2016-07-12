package node

import (
	"gopkg.in/mgo.v2/bson"
)

func ( this *NodeController ) Get_edit(){
	this.Data["page_title"] = "Node";
	this.Data["page_via_title"] = "Edit Form";

	this.TplName =  this.GetThemesAdmin() + "controllers/node/edit.html"
	this.Data["uniqid"] = this.Ctx.Input.Param(":Uniqid")
	node := this.get_Single()
	this.Data["node"] = node
}

func ( this *NodeController ) Post_edit(){

	form := EditForm{}

	if err := this.ParseForm( &form ); err != nil{
	}

	err := this.DB.C( "node" ).Update( bson.M{"uniqid": this.Ctx.Input.Param(":Uniqid")  },form )
	if err == nil{
		this.Data["json"] = map[string]interface{}{ "errorCode": 0, "msg":"Updated!" }
		this.ServeJSON()
		return
	}else{
		this.Data["json"] = map[string]interface{}{ "errorCode": 1, "msg":  "uniquid " + this.Ctx.Input.Param(":Uniqid")  + "Faild Update" }
		this.ServeJSON()
	}

}




/**
 * get single data
 */
func (  this *NodeController ) get_Single() Node{
	this.Uniqid = this.Ctx.Input.Param(":Uniqid")
	var result Node
	err := this.DB.C( "node" ).Find(bson.M{"uniqid":  this.Uniqid   }).One(&result)
	if err ==nil {
		return result
	}else{
		return Node{}
	}
}


