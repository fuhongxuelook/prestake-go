package Models

import (
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Stake struct {
    Id          int    `orm:"column(id)" json:"id"`
    Address     string `orm:"column(address);description(用户地址)" json:"address"`
    Amount      uint64 `orm:"column(amount);description(数量)" json:"amount"`
    Status      uint8  `orm:"column(status);description(状态)" json:"status"`
    CreatedAt   string `orm:"description(创建时间);column(created_at)" json:"created_at"`
}


type Ido struct {
    Id          int    `orm:"column(id)" json:"id"`
    Address     string `orm:"column(address);description(用户地址)" json:"address"`
    Amount      uint64 `orm:"column(amount);description(数量)" json:"amount"`
    Status      uint8  `orm:"column(status);description(状态)" json:"status"`
    CreatedAt   string `orm:"description(创建时间);column(created_at)" json:"created_at"`
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // set default database
    orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1)/skp?charset=utf8")
    orm.Debug = true
    // // // register model
    // orm.RegisterModel(new(Stu))
    orm.RegisterModel(new(Stake))
    orm.RegisterModel(new(Ido))

    orm.RunSyncdb("default", false, true)


    // // // create table
    // orm.RunSyncdb("mermaidnft", false, true)
}