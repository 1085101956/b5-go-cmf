package config

type Config struct {
	System   SystemConf              `yaml:"system"`
	Server   ServerConf              `yaml:"server"`
	Route    RouteConf               `yaml:"route"`
	DataBase map[string]DataBaseConf `yaml:"database"`
	Redis    map[string]RedisConf    `yaml:"redis"`
	Oss      OssConf                 `yaml:"oss"`
}
