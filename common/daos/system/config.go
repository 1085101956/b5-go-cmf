package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"sync"
)

type ConfigDao struct {
	Model *ConfigModel
}

var (
	instanceConfigDao *ConfigDao //单例的对象
	onceConfigDao     sync.Once
)

func NewConfigDao() *ConfigDao {
	onceConfigDao.Do(func() {
		instanceConfigDao = &ConfigDao{Model: NewConfigModel()}
	})
	return instanceConfigDao
}

func (d *ConfigDao) GetInfoByType(code string) *ConfigModel {
	if code == "" {
		return nil
	}
	info := d.Model.New()
	err := core.NewDao(d.Model).SetField("id,value,extra,style").First(info, "`type` = ?", code)
	if err != nil {
		return nil
	}
	return info
}

func (d *ConfigDao) GroupChildList() *[]ConfigModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).Lists(list, "`groups` != ''")
	return list
}
