package lib

// SetStatusData 用来后台公共操作设置状态
type SetStatusData struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Field  string `json:"field"`
}

// RoleMenuDataAuth 角色菜单和数据权限提交绑定的结构体
type RoleMenuDataAuth struct {
	Id        string `json:"id"`
	TreeId    string `json:"tree_id"`
	DataScope string `json:"data_scope"`
}
