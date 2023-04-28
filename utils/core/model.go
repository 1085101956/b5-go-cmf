package core

type IModel interface {
	Table() string    //返回表名
	INew() IModel     //创建一个新的model结构体实例
	GetId() string    //获取Id值
	DataBase() string // 表所在数据库
}

type MTotal struct {
	Total int64
}
