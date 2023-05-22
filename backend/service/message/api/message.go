package main

import (
	"flag"

	"tourBooking/service/message/api/internal/config"
	"tourBooking/service/message/api/internal/handler"
	"tourBooking/service/message/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/message-api.yaml", "the config file")

type MessageService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewMessageService(server *rest.Server) *MessageService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &MessageService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *MessageService) Start() error {
	return nil
}
