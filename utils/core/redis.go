// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

/////      redis的简单封装，可以支持多个redis     /////
/////      使用方法 G_Redis.Conn().Set     /////
/////      G_Redis.Conn("xxx") 可以传入redis标识前缀 默认为default     /////

package core

import (
	"github.com/go-redis/redis"
	"log"
)

type B5Redis struct {
	List map[string]*redis.Client
}

// InitRedis redis配置初始化
// 将多数据库进行sqlx.Open打开并赋值给全局的G_DB
func InitRedis() {
	db := &B5Redis{}
	db.parseItem()
	G_Redis = db
}

// Conn 数据库操作方法，返回*redis.Client
func (db *B5Redis) Conn(args ...string) *redis.Client {
	types := "default"
	if len(args) > 0 && args[0] != "" {
		types = args[0]
	}

	if val, exist := db.List[types]; exist {
		err := val.Ping().Err()
		if err != nil {
			log.Println("redis[" + types + "]连接失败：" + err.Error())
		}
		return val
	}
	log.Println("redis[" + types + "]不存在")
	return &redis.Client{}
}

// parseItem 解析并创建连接
func (db *B5Redis) parseItem() {
	list := make(map[string]*redis.Client)
	for key, item := range G_CONFIG.Redis {
		options := &redis.Options{
			Addr:         item.Host + ":" + item.Port,
			Password:     item.Password,
			MinIdleConns: 10,
		}
		client := redis.NewClient(options)
		list[key] = client
	}
	db.List = list
}
