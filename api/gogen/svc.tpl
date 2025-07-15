package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
	Trans     *i18n.Translator
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Trans:     trans,
		{{.middlewareAssignment}}
	}
}
