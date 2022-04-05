package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "prestake/models"
)

type IDOWLController struct {
	beego.Controller
}

func (c *IDOWLController) AddWl() {
	amount,_ := c.GetUint64("amount")
	address := c.GetString("address")

	Model.IDOReadAndCreateOrUpdate(address, amount)
	c.Data["json"] = true
    c.ServeJSON()
}


// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }