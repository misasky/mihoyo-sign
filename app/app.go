package main

import (
	"flag"
	"fmt"
	"mihoyo-sign/app/internal/cron"
	"mihoyo-sign/middleware"

	"mihoyo-sign/app/internal/config"
	"mihoyo-sign/app/internal/handler"
	"mihoyo-sign/app/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	if err := config.CheckAccountConf(); err != nil {
		panic(err)
	}
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()
	server.Use(middleware.NewCORS().SetCORS)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	cronJob := cron.NewCronJob()
	if err := cronJob.StartTask(); err != nil {
		panic(err)
	}
	defer cronJob.StopTask()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
