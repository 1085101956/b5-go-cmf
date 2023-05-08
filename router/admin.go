package router

import (
	"b5gocmf/admin/controller/books"
	"b5gocmf/admin/controller/common"
	"b5gocmf/admin/controller/demo"
	"b5gocmf/admin/controller/system"
	"b5gocmf/admin/lib"
	"b5gocmf/admin/middleware"
	"b5gocmf/utils/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (router *Router) Admin(engine *gin.Engine) {
	//加载管理端全局数据
	lib.AdminGlobalSetting(engine)

	//后端前缀
	adminPrefix := core.G_CONFIG.Route.Admin
	if adminPrefix == "" {
		adminPrefix = "/"
	}
	//路由分组 中间件
	var group *gin.RouterGroup
	group = engine.Group(adminPrefix, middleware.LoginAdminMiddleWare(), middleware.AuthAdminMiddleWare())

	//配置路由
	indexC := common.NewIndexController()
	//自动跳转到后台
	fmt.Println(core.G_CONFIG.Route.AutoAdmin)
	if core.G_CONFIG.Route.AutoAdmin {
		engine.GET(adminPrefix, func(ctx *gin.Context) {
			ctx.Redirect(http.StatusFound, indexC.ParseUrl("index", true))
		})
	}
	//common 分组
	indexC.Route(engine, group)
	common.NewPublicController().Route(engine, group)
	common.NewCommonController().Route(engine, group)

	books.NewBooksController().Route(engine, group)
	//system分组
	system.NewNoticeController().Route(engine, group)
	system.NewLoginLogController().Route(engine, group)
	system.NewAdminController().Route(engine, group)
	system.NewMenuController().Route(engine, group)
	system.NewConfigController().Route(engine, group)
	system.NewStructController().Route(engine, group)
	system.NewPositionController().Route(engine, group)
	system.NewRoleController().Route(engine, group)

	//demo分组
	demo.NewGenController().Route(engine, group)
	demo.NewBuildController().Route(engine, group)
	demo.NewMediaController().Route(engine, group)
	demo.NewTestInfoController().Route(engine, group)

	//__new_router__
}
