package system

import (
	"b5gocmf/utils/core"
	"sync"
)

type RoleMenuModel struct {
	RoleId string `db:"role_id" json:"role_id" form:"role_id" `
	MenuId string `db:"menu_id" json:"menu_id" form:"menu_id"`
}

func (m *RoleMenuModel) Table() string {
	return "b5net_role_menu"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *RoleMenuModel) INew() core.IModel {
	return m.New()
}

func (m *RoleMenuModel) GetId() string {
	return ""
}

func (m *RoleMenuModel) DataBase() string {
	return ""
}

func (m *RoleMenuModel) New() *RoleMenuModel {
	return &RoleMenuModel{}
}

func (m *RoleMenuModel) NewSlice() *[]RoleMenuModel {
	return &[]RoleMenuModel{}
}

var (
	instanceRoleMenuModel *RoleMenuModel //单例模式
	onceRoleMenuModel     sync.Once
)

// NewRoleMenuModel 单例获取config的结构体
func NewRoleMenuModel() *RoleMenuModel {
	onceRoleMenuModel.Do(func() {
		instanceRoleMenuModel = &RoleMenuModel{}
	})
	return instanceRoleMenuModel
}
