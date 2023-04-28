package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"b5gocmf/utils/types"
	"sync"
)

type MenuDao struct {
	Model *MenuModel
}

var (
	instanceMenuDao *MenuDao //单例的对象
	onceMenuDao     sync.Once
)

func NewMenuDao() *MenuDao {
	onceMenuDao.Do(func() {
		instanceMenuDao = &MenuDao{Model: NewMenuModel()}
	})
	return instanceMenuDao
}

func (d *MenuDao) MenuTreeList() *[]MenuModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetField("id,parent_id,name").SetOrderBy([]types.KeyVal{{Key: "parent_id", Value: "asc"}, {Key: "list_sort", Value: "asc"}, {Key: "id", Value: "asc"}}).Lists(list, "")
	return list
}
func (d *MenuDao) MenuLists() *[]MenuModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetOrderBy([]types.KeyVal{{Key: "parent_id"}, {Key: "list_sort"}, {Key: "id"}}).Lists(list, "")
	return list
}
func (d *MenuDao) GetMenuShowLists(idList []string) *[]MenuModel {
	list := d.Model.NewSlice()
	where := "`type` != ?"
	args := make([]any, 0)
	args = append(args, "F")
	if idList != nil {
		where += " AND id in (?)"
		args = append(args, idList)
	}
	_ = core.NewDao(d.Model).SetOrderBy([]types.KeyVal{{Key: "parent_id", Value: "asc"}, {Key: "list_sort", Value: "asc"}, {Key: "id", Value: "asc"}}).Lists(list, where, args...)
	return list
}
