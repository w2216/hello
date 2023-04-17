package logic

import (
	"context"
	"math/rand"

	"test/services/hello/api/internal/svc"
	"test/services/hello/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req *types.PostIndexReq) (resp *types.PostIndexResp, err error) {
	id := rand.Int63n(100)
	resp = &types.PostIndexResp{
		Id:   id,
		Name: req.Name,
	}

	return
}
