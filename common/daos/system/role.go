package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"b5gocmf/utils/types"
	"sync"
)

type RoleDao struct {
	Model *RoleModel
}

var (
	instanceRoleDao *RoleDao //单例的对象
	onceRoleDao     sync.Once
)

func NewRoleDao() *RoleDao {
	onceRoleDao.Do(func() {
		instanceRoleDao = &RoleDao{Model: NewRoleModel()}
	})
	return instanceRoleDao
}

func (d *RoleDao) GetInfoById(id string) *RoleModel {
	model := d.Model.New()
	err := core.NewDao(d.Model).First(model, id)
	if err != nil || model.Id == "" {
		return nil
	}
	return model
}

func (d *RoleDao) GetLists() *[]RoleModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetOrderBy([]types.KeyVal{{Key: "list_sort"}, {Key: "id"}}).Lists(list, "")
	return list
}
