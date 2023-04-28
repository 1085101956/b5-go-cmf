package system

import (
	. "b5gocmf/common/daos/system"
	"errors"
	"strings"
	"sync"
)

type AdminStructService struct {
	Dao *AdminStructDao
}

var (
	instanceAdminStructService *AdminStructService //单例的对象
	onceAdminStructService     sync.Once
)

func NewAdminStructService() *AdminStructService {
	onceAdminStructService.Do(func() {
		instanceAdminStructService = &AdminStructService{Dao: NewAdminStructDao()}
	})
	return instanceAdminStructService
}

// GetStructIdListByAdmin 获取某个角色的组织ID数组
func (s *AdminStructService) GetStructIdListByAdmin(adminId string) (reList []string) {
	list, err := s.Dao.GetStructListByAdmin(adminId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.StructId)
	}
	return
}

// GetAdminIdListByStruct 获取某个角色的组织ID数组
func (s *AdminStructService) GetAdminIdListByStruct(structId string) (reList []string) {
	list, err := s.Dao.GetAdminListByStruct(structId)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.AdminId)
	}
	return
}

func (s *AdminStructService) GetAdminIdListByStructList(structIds []string) (reList []string) {
	list, err := s.Dao.GetAdminListByStructList(structIds)
	if err != nil {
		return
	}
	for _, v := range *list {
		reList = append(reList, v.AdminId)
	}
	return
}

// UpdateAdminStruct 更新角色的组织节点
func (s *AdminStructService) UpdateAdminStruct(adminId string, structIds string) error {
	if adminId == "" {
		return nil
	}
	err := s.Dao.DeleteByAdminId(adminId)
	if err != nil {
		return errors.New("删除原节点失败")
	}
	if structIds == "" {
		return nil
	}
	s.Dao.InsertList(adminId, strings.Split(structIds, ","))
	return nil
}
