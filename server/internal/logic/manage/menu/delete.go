package menu

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/manage/menu"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}
	// whether it has submenu
	subMenus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "parent_id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	if err != nil {
		return
	}
	if len(subMenus) > 0 {
		return nil,
	}
	err = l.svcCtx.Model.ManageMenu.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	return
}
