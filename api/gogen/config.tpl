package config

import (
    "github.com/yyzyyyzy/syhr-common/casbin"
	"github.com/yyzyyyzy/syhr-common/cros"
	"github.com/yyzyyyzy/syhr-common/i18n"
	"github.com/yyzyyyzy/syhr-common/nacos"
	"github.com/yyzyyyzy/syhr-common/orm"
	"github.com/yyzyyyzy/syhr-common/redis"
    {{.authImport}}
)

type Config struct {
	rest.RestConf
	I18nConf     i18n.Conf
	DatabaseConf orm.Config
	AuthConf     rest.AuthConf
	CROSConf     cros.CROSConf
	NacosConf    nacos.NacosConf
	RedisConf    redis.RedisConf
	CasbinConf   casbin.CasbinConf
	{{.jwtTrans}}
}
