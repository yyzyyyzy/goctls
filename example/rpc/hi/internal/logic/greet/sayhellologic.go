package greetlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/yyzyyyzy/goctls/example/rpc/hi/internal/svc"
	"github.com/yyzyyyzy/goctls/example/rpc/hi/pb/hi"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *hi.HelloReq) (*hi.HelloResp, error) {
	// todo: add your logic here and delete this line

	return &hi.HelloResp{}, nil
}
