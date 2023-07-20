package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	InstanceId int
	DataSource string
	Auth       struct {
		AccessSecret string
	}
}
