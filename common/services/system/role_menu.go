package system

import (
	. "b5gocmf/common/daos/system"
	"errors"
	"strings"
	"sync"
)

type RoleMenuService struct {
	Dao *RoleMenuDao
}

var (
	instanceRoleMenuService *RoleMenuService //单例的对象
	onceRoleMenuService     sync.Once
)

func NewRoleMenuService() *RoleMenuService {
	onceRoleMenuService.Do(func() {
		instanceRoleMenuService = &RoleMenuService{Dao: NewRoleMenuDao()}
	})
	return instanceRoleMenuService
}

// GetRoleMenuIdList 获取某个角色的菜单ID数组
func (s *RoleMenuService) GetRoleMenuIdList(roleId string) (reList []string) {
	list, err := s.Dao.GetRoleMenuList(roleId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.MenuId)
	}
	return
}

// UpdateRoleMenu 更新角色的菜单节点
func (s *RoleMenuService) UpdateRoleMenu(roleId string, menuIds string) error {
	if roleId == "" {
		return nil
	}
	err := s.Dao.DeleteByRoleId(roleId)
	if err != nil {
		return errors.New("删除原节点失败")
	}
	if menuIds == "" {
		return nil
	}
	s.Dao.InsertList(roleId, strings.Split(menuIds, ","))
	return nil
}
