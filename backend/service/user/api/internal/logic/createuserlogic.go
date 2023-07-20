package logic

import (
	"context"
	"database/sql"

	"tourBooking/common"
	"tourBooking/service/user/api/internal/svc"
	"tourBooking/service/user/api/internal/types"
	"tourBooking/service/user/api/internal/utils"
	"tourBooking/service/user/model"

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
	currentTime := common.GetCurrentTime()

	if req == nil {
		l.Logger.Info(err)
		return &types.RegistationResp{
			Message: common.INVALID_REQUEST,
		}, nil
	}
	//check exist name
	checkName, err := l.svcCtx.UserModel.FindByUserName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Info(err)
		return &types.RegistationResp{
			Message: common.ERROR_DB,
		}, nil
	}
	if checkName != nil {
		l.Logger.Info(err)
		return &types.RegistationResp{
			Message: common.USER_EXIST,
		}, nil
	}
	hashPw := utils.HashPassword(req.Password)
	registation = &model.Users{
		UserId:    l.svcCtx.ObjSync.GenServiceObjID(),
		Username:  sql.NullString{String: req.Name, Valid: true},
		Email:     sql.NullString{String: req.Email, Valid: true},
		Password:  sql.NullString{String: hashPw, Valid: true},
		Gender:    sql.NullString{String: req.Gender, Valid: true},
		Dob:       sql.NullInt64{Int64: req.Dob, Valid: true},
		CreatedAt: sql.NullInt64{Int64: currentTime, Valid: true},
	}
	_, err = l.svcCtx.UserModel.Insertvalues(l.ctx, registation)
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
