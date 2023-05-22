package main

import (
	"flag"

	"tourBooking/service/payment/api/internal/config"
	"tourBooking/service/payment/api/internal/handler"
	"tourBooking/service/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/payment-api.yaml", "the config file")

type PaymentService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewPaymentService(server *rest.Server) *PaymentService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &PaymentService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *PaymentService) Start() error {
	return nil
}
