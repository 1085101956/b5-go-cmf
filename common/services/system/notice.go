package system

import (
	. "b5gocmf/common/daos/system"
	. "b5gocmf/utils/types"
	"sync"
)

type NoticeService struct {
	Dao *NoticeDao
}

var (
	instanceNoticeService *NoticeService //单例的对象
	onceNoticeService     sync.Once
)

func NewNoticeService() *NoticeService {
	onceNoticeService.Do(func() {
		instanceNoticeService = &NoticeService{Dao: NewNoticeDao()}
	})
	return instanceNoticeService
}

// TypeList 类型集合
func (s *NoticeService) TypeList() []KeyVal {
	return []KeyVal{{Key: "1", Value: "通知"}, {Key: "2", Value: "公告"}}
}

// GetTypeName 配置分组列表
func (s *NoticeService) GetTypeName(types string) string {
	name := ""
	if types != "" {
		typeList := s.TypeList()
		for _, v := range typeList {
			if v.Key == types {
				name = v.Value
				break
			}
		}
	}
	return name
}
