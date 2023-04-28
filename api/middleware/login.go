////    登录中间件 只对token进行解析，成功将信息放到上下文中    ////
////    app_token表存储token，并且可以区分多平台和用户类型    ////

package middleware

import (
	. "b5gocmf/common/daos/system"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HandleLogin struct {
	Plat     string
	UserType string
	IsLogin  bool
}

// ApiToken 只接收application/json请求中的token
type ApiToken struct {
	Token string `json:"token"`
}

func NewUserLoginMiddleWare() *HandleLogin {
	return &HandleLogin{Plat: "app", UserType: "user"}
}

func NewStoreLoginMiddleWare() *HandleLogin {
	return &HandleLogin{Plat: "app", UserType: "store"}
}

func (h *HandleLogin) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := &ApiToken{}
		_ = context.ShouldBindBodyWith(token, binding.JSON)

		if token.Token != "" {
			model := NewAppTokenDao().GetToken(token.Token, h.UserType, h.Plat)
			if model != nil && model.Token != "" {
				context.Set("_token_", model)
			}
		}
	}
}
