package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neocxf/oa/app/oa"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
)

// WechatRoutes func for describe group of private routes.
func WechatRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	//获取wechat实例
	wc := InitWechat()

	//公众号例子相关操作
	exampleOffAccount := oa.NewAccount(wc)
	//处理推送消息以及事件
	route.All("/serve", exampleOffAccount.Serve)
	//获取ak
	route.Get("/oa/basic/get_access_token", exampleOffAccount.GetAccessToken)
	//获取微信callback IP
	route.Get("/oa/basic/get_callback_ip", exampleOffAccount.GetCallbackIP)
	//获取微信API接口 IP
	route.Get("/oa/basic/get_api_domain_ip", exampleOffAccount.GetAPIDomainIP)
	//清理接口调用次数
	route.Get("/oa/basic/clear_quota", exampleOffAccount.ClearQuota)

	//获取
	route.Get("/oa/menu", exampleOffAccount.GetMenus)
	route.Post("/oa/menu", exampleOffAccount.NewMenu)
}

// InitWechat 获取wechat实例
// 在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	// cfg := config.GetConfig()
	// redisOpts := &cache.RedisOpts{
	// 	Host:        cfg.Redis.Host,
	// 	Password:    cfg.Redis.Password,
	// 	Database:    cfg.Redis.Database,
	// 	MaxActive:   cfg.Redis.MaxActive,
	// 	MaxIdle:     cfg.Redis.MaxIdle,
	// 	IdleTimeout: cfg.Redis.IdleTimeout,
	// }
	// redisCache := cache.NewRedis(redisOpts)
	// wc.SetCache(redisCache)

	memory := cache.NewMemory()
	wc.SetCache(memory)
	return wc
}
