package logic

import (
	"context"
	"database/sql"
	"time"

	"tourBooking/common"
	"tourBooking/service/user/api/internal/svc"
	"tourBooking/service/user/api/internal/types"
	"tourBooking/service/user/api/internal/utils"
	"tourBooking/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	l.Logger.Info("Login", req)
	var Result *model.Users
	if req == nil {
		l.Logger.Info(err)
		return &types.LoginResp{
			Message: common.INVALID_REQUEST,
		}, nil
	}
	Result, err = l.svcCtx.UserModel.FindByUserName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Info(err)
		return &types.LoginResp{
			Message: common.ERROR_DB,
		}, nil
	}

	token, err := utils.GenToken(req.Name+time.Now().String(), l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		l.Logger.Info(err)
		return &types.LoginResp{
			Message: common.ERROR_TOKEN_CREATED_FAIL,
		}, nil
	}

	pw := sql.NullString(Result.Password)
	bool, msg := utils.VerifyPassword(pw.String, req.Password)
	if !bool {
		l.Logger.Info("invalid pass")
		return &types.LoginResp{
			Message: msg,
		}, nil
	}
	return &types.LoginResp{
		Message: common.SUCCESSFUL,
		Token:   token,
	}, nil
}
