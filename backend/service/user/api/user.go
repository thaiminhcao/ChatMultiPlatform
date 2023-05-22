package api

import (
	"flag"

	"tourBooking/service/user/api/internal/config"
	"tourBooking/service/user/api/internal/handler"
	"tourBooking/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

type UserService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewUserService(server *rest.Server) *UserService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &UserService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (cs *UserService) Start() error {
	return nil
}
