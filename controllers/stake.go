package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "prestake/models"
)

type StakeController struct {
	beego.Controller
}

func (c *StakeController) Stake() {
	amount := c.GetString("amount")
	address := c.GetString("address")

	Model.ReadAndCreateOrUpdate(address, amount)
	c.Data["json"] = true
    c.ServeJSON()
}


func (c *StakeController) Rank() {
	amount := c.GetString("amount")
	address := c.GetString("address")
	
	rank := Model.GetAddressRand(address, amount)

	c.Data["json"] = rank
    c.ServeJSON()
}


// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }