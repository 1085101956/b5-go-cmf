package system

import (
	. "b5gocmf/common/models/system"
	"sync"
)

type NoticeDao struct {
	Model *NoticeModel
}

var (
	instanceNoticeDao *NoticeDao //单例的对象
	onceNoticeDao     sync.Once
)

func NewNoticeDao() *NoticeDao {
	onceNoticeDao.Do(func() {
		instanceNoticeDao = &NoticeDao{Model: NewNoticeModel()}
	})
	return instanceNoticeDao
}
