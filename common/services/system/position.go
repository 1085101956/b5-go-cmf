package system

import (
	. "b5gocmf/common/daos/system"
	"sync"
)

type PositionService struct {
	Dao *PositionDao
}

var (
	instancePositionService *PositionService //单例的对象
	oncePositionService     sync.Once
)

func NewPositionService() *PositionService {
	oncePositionService.Do(func() {
		instancePositionService = &PositionService{Dao: NewPositionDao()}
	})
	return instancePositionService
}

// GetName 获取岗位名称
func (s *PositionService) GetName(id string) string {
	if id == "0" || id == "" {
		return ""
	}
	info := s.Dao.GetInfoById(id)
	if info == nil {
		return ""
	}
	return info.Name
}
