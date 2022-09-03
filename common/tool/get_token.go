package tool

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"quora/app/user/model"
)

func GetUserFromCtx(ctx context.Context) model.MyClaims {
	var User model.MyClaims
	if jsonUid, ok := ctx.Value("uid").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			User.Id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidErr : %+v", err)
		}
	}

	if jsonUid, ok := ctx.Value("name").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			User.Id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetNameErr : %+v", err)
		}
	}

	return User
}
func JudgeLoginFormCtx(ctx context.Context) model.MyClaims {
	var User model.MyClaims
	if jsonUid, ok := ctx.Value("uid").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			User.Id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidErr : %+v", err)
		}
	}

	if jsonUid, ok := ctx.Value("name").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			User.Id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidErr : %+v", err)
		}
	}

	return User
}
