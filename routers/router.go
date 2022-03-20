package routers

import (
	"prestake/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/prestake/list", &controllers.ListController{},"get:List")
    beego.Router("/prestake/stake", &controllers.StakeController{},"get:Stake")
    beego.Router("/prestake/rank", &controllers.StakeController{},"get:Rank")
}
