package logic

import (
	"context"

	"go-zero-demo/book/service/search/api/internal/svc"
	"go-zero-demo/book/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	return &types.SearchReply{}, nil
}
