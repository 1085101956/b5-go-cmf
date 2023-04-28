package middleware

import (
	"b5gocmf/admin/lib"
	"b5gocmf/admin/services"
	"b5gocmf/common/services/system"
	"b5gocmf/utils/core"
	"b5gocmf/utils/tool"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthAdminMiddleWare 权限判断中间件
func AuthAdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		fmt.Println(ctx, core.G_CONFIG.Route.Admin)

		prefix := core.G_CONFIG.Route.Admin
		if prefix == path {
			return
		}
		auth := CheckAuth(ctx, path)
		if !auth {
			if tool.IsRender(ctx) {
				lib.ErrorHtml(ctx, errors.New("无权限访问"))
			} else {
				ctx.JSON(http.StatusOK, lib.JsonArgsParse(false, "无权限访问"))
			}
			ctx.Abort()
		}
	}
}

// CheckAuth 权限检测
func CheckAuth(ctx *gin.Context, path string) bool {
	maps := lib.AdminPathParse(path)
	if maps == nil {
		return false
	}
	group := maps["group"]
	controller := maps["controller"]
	action := maps["action"]

	if controller == "public" || controller == "common" || controller == "index" {
		return true
	}
	loginData := services.GetLoginByCtx(ctx)
	if loginData == nil {
		return false
	}
	if loginData.IsAdmin == "1" {
		return true
	}
	if len(loginData.MenuList) < 1 {
		return false
	}

	perms := group + ":" + controller + ":" + action
	perms = strings.Trim(perms, ":")

	menuId := system.NewMenuService().CheckPerms(perms)
	if menuId == "" {
		return false
	}

	menuIdList := "," + strings.Join(loginData.MenuList, ",") + ","
	if strings.Index(menuIdList, menuId) > -1 {
		return true
	}
	return false
}
