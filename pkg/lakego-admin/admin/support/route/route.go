package route

import (
    gin "github.com/deatil/lakego-admin/lakego/router"
    "github.com/deatil/lakego-admin/lakego/facade/config"
    "github.com/deatil/lakego-admin/lakego/facade/router"

    "github.com/deatil/lakego-admin/lakego/route"
)

// 路由
func AddRoute(f func(rg *gin.RouterGroup)) {
    // 路由
    engine := route.New().Get()

    // 配置
    conf := config.New("admin")

    // 后台路由及设置中间件
    m := router.GetMiddlewares(conf.GetString("Route.Middleware"))

    // 路由
    admin := engine.Group(conf.GetString("Route.Prefix"))
    {
        admin.Use(m...)
        {
            f(admin)
        }
    }
}

