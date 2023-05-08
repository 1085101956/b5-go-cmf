package books

import (
	"b5gocmf/utils/core"
	"sync"
	"time"
)

type BooksModel struct {
	Id          string     `json:"id" form:"id" `
	Img         string     `json:"img" form:"img"`
	StructId    string     `db:"struct_id" json:"struct_id" form:"-"` // 组织ID
	UserId      string     `db:"user_id" json:"user_id" form:"-"`     // 用户ID
	Imgs        string     `json:"imgs" form:"imgs"`
	Crop        string     `json:"crop" form:"crop" `
	Video       string     `json:"video" form:"video"`
	File        string     `json:"file" form:"file"`
	Files       string     `json:"files" form:"files"`
	CreateTime  *time.Time `db:"create_time" json:"create_time" form:"-"`
	UpdateaTime *time.Time `db:"update_time" json:"update_time" form:"-"`
}

func (m *BooksModel) Table() string {
	return "b5net_books"
}

func (m *BooksModel) INew() core.IModel {
	return m.New()
}
func (m *BooksModel) GetId() string {
	return m.Id
}

func (m *BooksModel) DataBase() string {
	return ""
}

func (m *BooksModel) New() *BooksModel {
	return &BooksModel{}
}

func (m *BooksModel) NewSlice() *[]BooksModel {
	return &[]BooksModel{}
}

var (
	instanceBooksModel *BooksModel //单例模式
	onceMediaModel     sync.Once
)

// NewBooksModel 单例获取config的结构体
func NewBooksModel() *BooksModel {
	onceMediaModel.Do(func() {
		instanceBooksModel = &BooksModel{}
	})
	return instanceBooksModel
}
