package config

type RouteConf struct {
	Admin     string `yaml:"admin"`      //后台路由访问前缀 最后不能带有/
	Api       string `yaml:"api"`        //接口访问前缀 最后不能带有/
	AutoAdmin bool   `yaml:"auto-admin"` //当访问后端路由前缀是否自动跳转到后台
}
