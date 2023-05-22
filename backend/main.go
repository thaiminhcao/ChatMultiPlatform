package main

import (
	"flag"
	"fmt"

	"tourBooking/config"
	userAPI "tourBooking/service/user/api"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("server_config", "etc/server.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//user service
	userService := userAPI.NewUserService(server)
	userService.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
