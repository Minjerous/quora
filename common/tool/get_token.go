package tool

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	jwtModel "quora/common/jwt"
)

func GetUserFromCtx(ctx context.Context) jwtModel.MyClaims {
	var User jwtModel.MyClaims
	if jsonUid, ok := ctx.Value("uid").(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			User.Id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidErr : %+v", err)
		}
	}

	if Name, ok := ctx.Value("name").(json.Number); ok {
		User.Name = Name.String()
	} else {
		logx.WithContext(ctx).Errorf("GetUidErr : %+v", errors.New("name  is null"))
	}

	return User
}

func JudgeLoginFormCtx(ctx context.Context) jwtModel.MyClaims {
	var User jwtModel.MyClaims
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
