package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
	Enforcer  *casbin.Enforcer
	Trans     *i18n.Translator
}

func NewServiceContext(c {{.config}}) *ServiceContext {
    redisClient := c.RedisConf.MustNewRedis()
    translator := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)
    enforcer := c.CasbinConf.MustNewEnforcer(c.Database.Type, c.DBConf.GetMysqlDSN(), c.RedisConf)
	return &ServiceContext{
		Config: c,
		Enforcer:  enforcer,
		Trans:     translator,
		{{.middlewareAssignment}}
	}
}
