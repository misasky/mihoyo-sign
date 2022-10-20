package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logx"
	"mihoyo-sign/app/internal/config"
	"mihoyo-sign/thirdparty/mihoyo"
)

func runSign(job *cron.Cron) error {
	if err := job.AddFunc("0 0 0/6 * * ?", func() {
		conf, err := config.GetAccountConf()
		if err != nil {
			panic(err)
		}
		for _, a := range conf.Accounts {
			logx.Infof("当前签到 uid [%s], cookie: %s", a.Uid, a.Cookie)
			if a.Cookie == "" {
				logx.Infof("uid [%s] cookie未设置", a.Uid)
				continue
			}
			api := mihoyo.NewGenShin(a.Cookie)
			info, err := api.GetSignInfo(a.Uid)
			if err != nil {
				logx.Errorf("get sign info [%s] got error %s", a.Uid, err.Error())
				return
			}
			if info.Data.IsSign {
				logx.Infof("uid [%s] 已经签到过了", a.Uid)
				continue
			}
			if err := api.Sign(); err != nil {
				logx.Errorf("sign [%s] got error %s", a.Uid, err.Error())
				return
			}
		}

	}); err != nil {
		return fmt.Errorf("run sign task got error: %s", err.Error())
	}
	return nil
}
