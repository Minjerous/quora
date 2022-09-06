package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"quora/app/mq/asynq/job/internal/config"
	"quora/app/mq/asynq/job/internal/logic"
	"quora/app/mq/asynq/job/internal/svc"
)

var configFile = flag.String("f", "etc/asynq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	if err := c.SetUp(); err != nil {
		logx.Infof("init asynq-job service failed err :%v", err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	cronJob := logic.NewProcessorJob(ctx, svcContext)

	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("asynq is err :%+v", err)
		os.Exit(1)
	}
}
