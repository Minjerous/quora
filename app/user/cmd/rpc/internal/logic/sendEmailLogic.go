package logic

import (
	"context"
	"fmt"
	"quora/app/user/cmd/rpc/internal/config"
	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"
	"quora/app/user/model"
	"quora/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *pb.SendEmailReq) (*pb.SendEmailResp, error) {
	if !tool.VerifyEmailFormat(in.Email) {
		return nil, ErrFormatError
	}
	rs := model.NewRedisStore(fmt.Sprintf("Email:%s", in.Email), l.svcCtx.Redis, 120, l.ctx)
	code := tool.RandCode()

	err := rs.SetRedisValue(code)
	err = tool.SendCodeByEmail(in.Email, code, config.Config{}.Email)
	if err != nil {
		return &pb.SendEmailResp{
			OK: false,
		}, err
	}

	fmt.Println(code)
	return &pb.SendEmailResp{
		OK: true,
	}, nil
}
func AsyncAdd(run func() error) {
	//TODO: 扔进异步协程池
	go run()
}
