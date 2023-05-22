package main

import (
	"flag"

	"tourBooking/service/tour/api/internal/config"
	"tourBooking/service/tour/api/internal/handler"
	"tourBooking/service/tour/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/tour-api.yaml", "the config file")

type TourService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewTourService(server *rest.Server) *TourService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &TourService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *TourService) Start() error {
	return nil
}
