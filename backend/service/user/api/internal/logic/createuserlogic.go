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

	"github.com/guregu/null"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.RegistationReq) (resp *types.RegistationResp, err error) {
	l.Logger.Infof("CreateUser", req)
	var registation *model.Users
	var dob null.Time

	if req == nil {
		l.Logger.Info(err)
		return &types.RegistationResp{
			Message: common.INVALID_REQUEST,
		}, nil
	}
	t, err := time.Parse("2-1-2006", req.Dob)
	if err != nil {
		l.Logger.Info("Invalid type for field dob")
		return &types.RegistationResp{
			Message: common.INVALID_REQUEST,
		}, nil
	}
	dob = null.TimeFrom(t)
	hashPw := utils.HashPassword(req.Password)
	registation = &model.Users{
		Username: sql.NullString{String: req.Name, Valid: true},
		Email:    sql.NullString{String: req.Email, Valid: true},
		Password: sql.NullString{String: hashPw, Valid: true},
		Gender:   sql.NullString{String: req.Gender, Valid: true},
		Dob:      dob.NullTime,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, registation)
	if err != nil {
		l.Logger.Info(err)
		return &types.RegistationResp{
			Message: common.ERROR_DB,
		}, nil
	}
	return &types.RegistationResp{
		Message: common.SUCCESSFUL,
	}, nil
}
