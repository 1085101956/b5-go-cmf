package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"b5gocmf/utils/tool"
	"b5gocmf/utils/trans"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"sync"
)

type LoginLogDao struct {
	Model *LoginLogModel
}

var (
	instanceLoginLogDao *LoginLogDao //单例的对象
	onceLoginLogDao     sync.Once
)

func NewLoginLogDao() *LoginLogDao {
	onceLoginLogDao.Do(func() {
		instanceLoginLogDao = &LoginLogDao{Model: NewLoginLogModel()}
	})
	return instanceLoginLogDao
}

// Trash 清空表
func (d *LoginLogDao) Trash() error {
	_, err := core.NewDao(d.Model).Exec("TRUNCATE TABLE " + d.Model.Table())
	if err != nil {
		return err
	}
	return nil
}

func (d *LoginLogDao) AddLog(ctx *gin.Context, username string, status string, msg string) {
	model := d.Model.New()
	model.Status = status
	model.Msg = msg
	model.LoginName = username
	ip := ctx.ClientIP()
	location := ""
	net := ""
	if ip == "::1" {
		ip = "127.0.0.1"
		location = "本机地址"
	}
	if location == "" && ip != "" {
		loc := tool.NewIpLoc()
		if loc != nil {
			info, err := loc.GetInfo(ip)
			if err == nil {
				location = info.Country + " " + info.Region + " " + info.City
				net = info.Isp
			}
		}
	}

	ua := user_agent.New(ctx.Request.UserAgent())
	browser, version := ua.Browser()
	model.Os = ua.Platform() + " " + ua.OS()
	model.Browser = browser + " " + version
	model.IpAddr = ip
	model.LoginLocation = location
	model.Net = net
	_, _ = core.NewDao(d.Model).Insert(trans.GetStructDBKV(model, []string{"id"}, true))
}
