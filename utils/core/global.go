package core

// ** 全局参数 ** //

import (
	"b5gocmf/config"
	"b5gocmf/utils/plugin/singleflight"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	G_CONFIG    *config.Config //配置文件
	G_DB        *B5Db
	G_Redis     *B5Redis
	G_TIME      string = "2006-01-02 15:04:05"
	G_DATE      string = "2006-01-02"
	G_GENID     *Node
	G_Validate  *ValidateCtx
	G_STRUCT_ID string                        = "100"                           //根组织ID
	G_ROLE_ID   string                        = "1"                             //超管角色ID
	G_ADMIN_ID  string                        = "10000"                         //超管ID
	G_Cache     *cache.Cache                  = cache.New(0, 2*time.Hour)       //全局缓存数据
	G_Single    *singleflight.OnceFlightGroup = &singleflight.OnceFlightGroup{} //防击穿
)
