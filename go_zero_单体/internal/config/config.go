package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"mail/gormc"
)

type Config struct {
	rest.RestConf
	//Mysql struct {
	//	DataSource string
	//}
	Mysql      gormc.Mysql
	CacheRedis cache.CacheConf
}
