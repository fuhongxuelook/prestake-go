package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "prestake/models"
)

type ListController struct {
	beego.Controller
}

const PageNum = 50

func (c *ListController) List() {
	page, _ := c.GetInt("page")

	skip := 0;
	if page > 1 {
		skip = page * PageNum;
	}
	//services.DoPairsRegister()
	list := Model.GetList(PageNum, skip)

	c.Data["json"] = list
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }