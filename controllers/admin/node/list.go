package node

import (
	"html/template"
	"gopkg.in/mgo.v2/bson"
)

type Stats struct{
	count string
}

func (this *NodeController) Get_list(){
	this.IsLogin()
	this.Data["page_title"] = "Node";
	this.Data["page_via_title"] = "Add Form";
	this.Data["xsrfdata"]= template.HTML(this.BaseController.XSRFFormHTML())
	this.TplName =  this.GetThemesAdmin() + "controllers/node/list.html"
	this.SetSkip()


	count , _ := this.DB.C( "node" ).Count()
	var result []Node
	this.DB.C( "node" ).Find(bson.M{}).Limit( this.BaseController.Pagenums ).Skip( this.Skip ).All( &result )

	this.Data["List"] = result
	this.Data["Count"] = count


	this.SetPaginator( this.BaseController.Pagenums, int64( count ) )




	this.LayoutSections["Paginator"] = this.GetThemesAdmin() + "paginator.html"

}


