package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"sync"
)

type AdminDao struct {
	Model *AdminModel
}

var (
	instanceAdminDao *AdminDao //单例的对象
	onceAdminDao     sync.Once
)

func NewAdminDao() *AdminDao {
	onceAdminDao.Do(func() {
		instanceAdminDao = &AdminDao{Model: NewAdminModel()}
	})
	return instanceAdminDao
}

func (d *AdminDao) GetInfoById(id string) *AdminModel {
	model := d.Model.New()
	err := core.NewDao(d.Model).First(model, id)
	if err != nil || model.Id == "" {
		return nil
	}
	return model
}
