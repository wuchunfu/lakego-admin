package provider

import (
    "github.com/deatil/lakego-doak/lakego/router"
    "github.com/deatil/lakego-doak/lakego/provider"

    admin_route "github.com/deatil/lakego-doak-admin/admin/support/route"

    "github.com/deatil/lakego-doak-extension/extension/extension"
    ext_router "github.com/deatil/lakego-doak-extension/extension/route"
    ext_middleware "github.com/deatil/lakego-doak-extension/extension/middleware"
)

// 路由中间件
var routeMiddlewares = map[string]router.HandlerFunc{
    // 操作日志
    "lakego-admin.extension": ext_middleware.NewBoot(),
}

// 中间件分组
var middlewareGroups = map[string][]string{
    // 扩展
    "lakego-admin": {
        "lakego-admin.extension",
    },
}

/**
 * 服务提供者
 *
 * @create 2023-4-19
 * @author deatil
 */
type Extension struct {
    provider.ServiceProvider
}

// 注册
func (this *Extension) Register() {
    // 中间件
    this.loadMiddleware()
}

// 引导
func (this *Extension) Boot() {
    // 路由
    this.loadRoute()

    // 加载扩展
    this.loadExtension()
}

// 导入中间件
func (this *Extension) loadMiddleware() {
    m := router.InstanceMiddleware()

    // 导入中间件
    for name, value := range routeMiddlewares {
        m.AliasMiddleware(name, value)
    }

    // 导入中间件分组
    for groupName, middlewares := range middlewareGroups {
        for _, middleware := range middlewares {
            m.PushMiddlewareToGroup(groupName, middleware)
        }
    }
}

// 导入路由
func (this *Extension) loadRoute() {
    // 后台路由
    admin_route.AddRoute(func(engine *router.RouterGroup) {
        ext_router.Route(engine)
    })
}

// 导入扩展
func (this *Extension) loadExtension() {
    m := extension.GetManager()

    m.CallBooting()

    // 加载扩展
    m.BootExtension(this.GetApp())

    m.CallBooted()
}

