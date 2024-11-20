package logic

import (
	"context"

	"foxit-otel-go/api/internal/svc"
	"foxit-otel-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type RegistryReq struct {
	RegistryGroup string `json:"registryGroup"`
	RegistryKey   string `json:"registryKey"`
	RegistryValue string `json:"registryValue"`
}

func (l *DemoLogic) Demo() (resp *types.DemoResp, err error) {

	resp = &types.DemoResp{
		Name: "yt",
		Age:  18,
	}
	return
}
