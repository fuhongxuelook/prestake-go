package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "prestake/models"
)

type ListController struct {
	beego.Controller
}

type Res struct{
	List []Model.Stake `json:"list"`
	Total int64 `json:"total"`
	CurrentPage int `json:"current_page"`
}

const PageNum = 50

func (c *ListController) List() {
	page, _ := c.GetInt("page")

	skip := 0;
	if page > 1 {
		skip = page * PageNum;
	}

	total := Model.GetListNum()
	//services.DoPairsRegister()
	list := Model.GetList(PageNum, skip)

	res := Res{
		CurrentPage : page,
		List: list,
		Total: total,
	}

	c.Data["json"] = &res
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }