package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	I18nConf     i18n.Conf
	DBConf       orm.Config
	AuthConf     rest.AuthConf
	CROSConf     cros.CROSConf
	NacosConf    nacos.NacosConf
	RedisConf    redis.RedisConf
	CasbinConf   casbin.CasbinConf
	{{.auth}}
	{{.jwtTrans}}
}
