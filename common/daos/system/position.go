package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"b5gocmf/utils/types"
	"sync"
)

type PositionDao struct {
	Model *PositionModel
}

var (
	instancePositionDao *PositionDao //单例的对象
	oncePositionDao     sync.Once
)

func NewPositionDao() *PositionDao {
	oncePositionDao.Do(func() {
		instancePositionDao = &PositionDao{Model: NewPositionModel()}
	})
	return instancePositionDao
}

func (d *PositionDao) GetLists() *[]PositionModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetOrderBy([]types.KeyVal{{Key: "list_sort"}, {Key: "id"}}).Lists(list, "")
	return list
}

func (d *PositionDao) GetInfoById(id string) *PositionModel {
	model := d.Model.New()
	err := core.NewDao(d.Model).First(model, id)
	if err != nil || model.Id == "" {
		return nil
	}
	return model
}
