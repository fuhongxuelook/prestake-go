package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "prestake/models"
)

type IDOListController struct {
	beego.Controller
}

type IDORes struct{
	List []Model.Ido `json:"list"`
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
}

const IDOPageNum = 50

func (c *IDOListController) List() {
	page, _ := c.GetInt("page")

	skip := 0;

	if page <= 0 {
		page = 1
	}

	skip = (page - 1) * IDOPageNum;

	total := int(Model.GetIDOListNum())
	// totalPage := 0
	// if total % PageNum > 0 {
	// 	totalPage = total / PageNum + 1
	// } else {
	// 	totalPage = total / PageNum
	// }
	//services.DoPairsRegister()
	list := Model.GetIDOList(IDOPageNum, skip)

	res := IDORes{
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