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
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
}

const PageNum = 50

func (c *ListController) List() {
	page, _ := c.GetInt("page")

	skip := 0;
	if page > 1 {
		skip = page * PageNum;
	}

	total := int(Model.GetListNum())
	totalPage := 0
	if total % PageNum > 0 {
		totalPage = total / PageNum + 1
	} else {
		totalPage = total / PageNum
	}
	//services.DoPairsRegister()
	list := Model.GetList(PageNum, skip)

	num := len(list)
	if num > 0 {
		for i := 0; i < num; i++ {
			length := len(list[i].Amount);
            list[i].Amount = list[i].Amount[:length - 9]
        }
	} 

	res := Res{
		CurrentPage : page,
		List: list,
		Total: totalPage,
	}

	c.Data["json"] = &res
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }