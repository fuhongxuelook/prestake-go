package routers

import (
	"prestake/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/list", &controllers.ListController{},"get:List")
    beego.Router("/stake", &controllers.StakeController{},"get:Stake")
    beego.Router("/rank", &controllers.StakeController{},"get:Rank")
}
