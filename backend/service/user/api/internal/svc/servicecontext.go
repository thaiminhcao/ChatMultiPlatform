package svc

import (
	"tourBooking/service/user/api/internal/config"
	"tourBooking/service/user/model"
	"tourBooking/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Cors      rest.Middleware
	ObjSync   sync.ObjSync
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource)),
		ObjSync:   *sync.NewObjSync(nil, c.InstanceId),
	}
}
