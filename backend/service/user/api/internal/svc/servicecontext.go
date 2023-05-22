package svc

import (
	"tourBooking/service/user/api/internal/config"
	"tourBooking/service/user/api/internal/middleware"
	"tourBooking/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Cors      rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource)),
		Cors:      middleware.NewCorsMiddleware().Handle,
	}
}
