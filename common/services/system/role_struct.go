package system

import (
	. "b5gocmf/common/daos/system"
	"errors"
	"strings"
	"sync"
)

type RoleStructService struct {
	Dao *RoleStructDao
}

var (
	instanceRoleStructService *RoleStructService //单例的对象
	onceRoleStructService     sync.Once
)

func NewRoleStructService() *RoleStructService {
	onceRoleStructService.Do(func() {
		instanceRoleStructService = &RoleStructService{Dao: NewRoleStructDao()}
	})
	return instanceRoleStructService
}

// GetRoleStructIdList 获取某个角色的组织ID数组
func (s *RoleStructService) GetRoleStructIdList(roleId string) (reList []string) {
	list, err := s.Dao.GetRoleStructList(roleId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.StructId)
	}
	return
}

// UpdateRoleStruct 更新角色的组织节点
func (s *RoleStructService) UpdateRoleStruct(roleId string, structIds string) error {
	if roleId == "" {
		return nil
	}
	err := s.Dao.DeleteByRoleId(roleId)
	if err != nil {
		return errors.New("删除原节点失败")
	}
	if structIds == "" {
		return nil
	}
	s.Dao.InsertList(roleId, strings.Split(structIds, ","))
	return nil
}
