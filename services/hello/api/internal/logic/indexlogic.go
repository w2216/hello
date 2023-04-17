package logic

import (
	"context"
	"math/rand"
	"strconv"

	"test/services/hello/api/internal/svc"
	"test/services/hello/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.GetIndexReq) (resp *types.GetIndexResp, err error) {
	id := rand.Int63n(10000)
	resp = &types.GetIndexResp{
		Msg: "你好游客: " + strconv.FormatInt(id, 10),
	}
	return
}
