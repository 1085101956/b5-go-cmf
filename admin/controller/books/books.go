package books

import (
	"b5gocmf/admin/lib"
	"b5gocmf/common/models/books"
	"b5gocmf/common/services/system"
	"b5gocmf/utils/trans"
	"github.com/gin-gonic/gin"
)

type BooksController struct {
	lib.Controller
}

func (c *BooksController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
}

func NewBooksController() *BooksController {
	c := &BooksController{}
	c.Id = "index"
	c.Group = "books"

	c.IModel = books.NewBooksModel()

	c.HookAction.IndexAfter = c.indexAfter
	c.HookAction.EditRender = c.EditRender
	c.HookAction.SaveBefore = c.SaveBefore
	c.HookAction.DropBefore = c.DropBefore
	return c
}

func (c *BooksController) Index(ctx *gin.Context) {
	c.Render(ctx, "index", nil)
}
func (c *BooksController) indexAfter(ctx *gin.Context, list []any) []any {
	reList := make([]any, 0)
	for _, item := range list {
		info := item.(books.BooksModel)
		model := trans.StructToMap(info, true)
		model["struct_name"] = system.NewStructService().GetName(info.Id, false)
		model["user_name"] = system.NewAdminService().GetName(info.UserId)
		reList = append(reList, model)
	}
	return reList
}
