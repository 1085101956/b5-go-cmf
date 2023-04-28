package services

import (
	. "b5gocmf/common/models/system"
	"github.com/gin-gonic/gin"
)

// GetApiLoginInfo 获取登录信息
func GetApiLoginInfo(ctx *gin.Context) *AppTokenModel {
	appToken, exists := ctx.Get("_token_")
	if !exists {
		return nil
	}
	switch appToken.(type) {
	case *AppTokenModel:
		model := appToken.(*AppTokenModel)
		if model.UserId == "" {
			return nil
		}
		return model
	}
	return nil
}
